# keycloak-commander
keycloak-commander is a work in progress. Its intent is to provide an easy to use CLI for Red Hat's Keycloak authentication service.
## Use
### Create config file
> This file will store your username and password in plain text at the time of writing this!

The config file is created at `~/.keycloak-commander.json`
```
Usage:
  keycloak-commander init [flags]

Flags:
  -h, --help              help for init
  -i, --host string       Hostname of the Keycloak server
  -p, --password string   Admin's password
  -r, --realm string      Realm to use
  -u, --username string   Admin's username
```
### Create a client
```
Usage:
  keycloak-commander create client [flags]

Each argument shall be a string representing a client name. At least one client name is required.

Example: keycloak-commander create client my-client another-client a-third-client
```
### Update a client
The filename of a json file representing a [ClientRepresentation](https://www.keycloak.org/docs-api/11.0/javadocs/org/keycloak/representations/idm/ClientRepresentation.html) needs to be passed into the program .

An example can be found at [`examples/update_client.json`](examples/update_client.json)
```
The payload should reflect the Client struct as defined here: https://github.com/Nerzal/gocloak/blob/fe4f627eaf1bff988ee5df2fd0d0b87daac6c074/models.go#L435

Example: keycloak-commander update client --json /path/to/file.json

Usage:
  keycloak-commander update client [flags]

Flags:
  -h, --help          help for client
      --json string   Json file representing the ClientRepresentation payload.
```
