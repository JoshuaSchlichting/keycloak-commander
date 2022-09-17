package keycloak

import (
	"testing"
)

func TestKeycloakCreateClient(t *testing.T) {
	keycloakCommander := NewKeycloakCommander("http://localhost:8081/", "admin", "admin", "master")
	// keycloakCommander.CreateClient("testdeletethis")
}
