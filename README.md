<p align="center">
  <h1 align="center">VaultSync</h1>
  <p align="center">Sync secrets from various platforms using plugins.</p>
</p>



## Description

VaultSync copies a platform's secrets to a specified instance of [HashiCorp Vault](https://vaultproject.io).

## Plugins

Each plugin defines the following:
- Access to the platform's API
- Path to the platform's secrets
- Desired structure for that platform's secrets in Vault

## Why?

It's important to establish a standard for maintaining external Vault secrets. If you rotate a vendor platform's API keys--such as after a security incident--then mitigation should not completely destroy your applications.

## Official Plugins

Currently none.

## Community Plugins

## Planned Plugins

| Name                                                       | Author                                                                                       |
| :--------------------------------------------------------- | :------------------------------------------------------------------------------------------- |
| Datadog                                                    | [@particledecay](https://github.com/particledecay)                                           |
