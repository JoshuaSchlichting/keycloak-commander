# keycloak-commander
keycloak-commander is a work in progress. Its intent is to provide and easy to use CLI for Red Hat's Keycloak authentication service.
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


Usage:
  keycloak-commander create client [flags]

Each argument shall be a string represent a client name. At least one client name is required.

Example: keycloak-commander create client my-client another-client a-third-client
```

