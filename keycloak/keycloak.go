package keycloak

import (
	"context"
	"crypto/tls"
	"log"

	gocloak "github.com/Nerzal/gocloak/v11"
)

type KeycloakCommander struct {
	Host        string
	Realm       string
	Username    string
	Password    string
	accessToken string
	context     context.Context
	client      gocloak.GoCloak
}

func NewKeycloakCommander(hostname, username, password, realm string) *KeycloakCommander {
	client := gocloak.NewClient(hostname, gocloak.SetAuthAdminRealms("admin/realms"), gocloak.SetAuthRealms("realms"))
	ctx := context.Background()
	restyClient := client.RestyClient()
	restyClient.SetDebug(false)
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	token, err := client.LoginAdmin(ctx, username, password, realm)
	if err != nil {
		log.Panicf("Something wrong with the credentials or url: %v", err)
	}
	return &KeycloakCommander{
		Host:        hostname,
		Realm:       realm,
		Username:    username,
		Password:    password,
		accessToken: token.AccessToken,
		context:     ctx,
		client:      client,
	}
}

func (kc *KeycloakCommander) CreateClient(clientName string) error {

	client, err := kc.client.CreateClient(kc.context, kc.accessToken, kc.Realm, gocloak.Client{
		ClientID: &clientName,
	})
	if err != nil {
		return err
	}
	log.Printf("Created new client with ID: %s\n", client)
	return nil
}

func (kc *KeycloakCommander) UpdateClient(updatedClient *gocloak.Client) error {

	err := kc.client.UpdateClient(kc.context, kc.accessToken, kc.Realm, *updatedClient)
	if err != nil {
		return err
	}
	log.Printf("Client updated: %s\n", *updatedClient.ClientID)
	return nil
}

func (kc *KeycloakCommander) CreateRealm(realmRepresentation *gocloak.RealmRepresentation) error {
	realm, err := kc.client.CreateRealm(kc.context, kc.accessToken, *realmRepresentation)
	if err != nil {
		return err
	}
	log.Printf("Created new realm with ID: %s\n", realm)
	return nil
}
