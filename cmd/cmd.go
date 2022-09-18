package cmd

import (
	"keycloak-commander/keycloak"
)

var KeycloakCommander *keycloak.KeycloakCommander

var GetFilePayload func(filename string) []byte
