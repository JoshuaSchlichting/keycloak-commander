package cmd

import (
	"keycloak-commander/keycloak"
)

var NewKeycloakCommander func(hostname, username, password, realm string) *keycloak.KeycloakCommander

var GetFilePayload func(filename string) []byte

var ConfigFileWriter func(data []byte) error

var ConfigPayload Config
