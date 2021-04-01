package vault

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

var client *api.Logical

// InitClient initializes a new connection to the Vault instance
func InitClient(vaultURL, token string) error {
	config := &api.Config{
		Address: vaultURL,
	}
	c, err := api.NewClient(config)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	c.SetToken(token)
	client = c.Logical()

	return nil
}

// GetSecret retrieves a secret from a Vault instance
func GetSecret(keyPath, fieldName string) (string, error) {
	secret, err := client.Read(keyPath)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	value, ok := secret.Data[fieldName].(string)
	if !ok {
		return "", fmt.Errorf("'%s' not found in '%s'", fieldName, keyPath)
	}

	return value, nil
}
