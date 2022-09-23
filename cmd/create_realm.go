package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Nerzal/gocloak/v11"
	"github.com/spf13/cobra"
)

// createRealmCmd represents the createRealm command
var createRealmCmd = &cobra.Command{
	Use:   "realm",
	Short: "Create a Keycloak realm",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 0 {
			log.Fatalf("0 arguments expected, got %d", len(args))
		}
		jsonFilename := cmd.Flag("json").Value.String()
		filePayload, err := os.ReadFile(jsonFilename)
		if err != nil {
			panic(err)
		}
		realmRepresentation := &gocloak.RealmRepresentation{}

		err = json.Unmarshal(filePayload, realmRepresentation)
		if err != nil {
			log.Fatal("There was an error loading the JSON file as a ClientRepresentation: ", err)
		}
		initKeycloakCommander()
		err = keycloakCommander.CreateRealm(realmRepresentation)
		if err != nil {
			log.Fatal("There was an error updating the client: ", err)
		}
	},
}

func init() {
	createCmd.AddCommand(createRealmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createRealmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createRealmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
