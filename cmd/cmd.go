package cmd

import (
	"keycloak-commander/keycloak"
)

var keycloakCommander *keycloak.KeycloakCommander

type Config struct {
	Host     string `json:"host"`
	Realm    string `json:"realm"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func initKeycloakCommander() {
	// load config
	keycloakCommander = keycloak.NewKeycloakCommander(
		ConfigPayload.Host,
		ConfigPayload.Username,
		ConfigPayload.Password,
		ConfigPayload.Realm,
	)
}
