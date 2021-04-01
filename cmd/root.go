package cmd

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	verbose    bool
	output     string
	vaultURL   string
	vaultToken string
)

var rootCmd = &cobra.Command{
	Use:   "vaultsync",
	Short: "A plugin-based tool for syncing secrets to Vault",
	Long: `VaultSync uses vendor-specific plugins to sync secrets to an
		   instance of HashiCorp Vault.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if output == "text" {
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		}
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

		// debug mode
		if verbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
		log.Debug().Msg("debug messaging turned on")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("an action is required")
		}
		return nil
	},
}

func init() {
	// use env vars as default values
	if addr, ok := os.LookupEnv("VAULT_ADDR"); ok {
		vaultURL = addr
	}
	if token, ok := os.LookupEnv("VAULT_TOKEN"); ok {
		vaultToken = token
	}

	// flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "display debug messages")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "text", "format of logging output ('text', 'json')")
	rootCmd.PersistentFlags().StringVarP(&vaultURL, "vault-addr", "", vaultURL, "address of Vault instance")
	rootCmd.PersistentFlags().StringVarP(&vaultToken, "vault-token", "", vaultToken, "token used to authenticate to Vault")
}

// Execute combines all of the available command functions
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Msgf("%v", err)
	}
}
