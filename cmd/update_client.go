package cmd

import (
	"encoding/json"
	"os"

	"github.com/Nerzal/gocloak/v11"
	"github.com/spf13/cobra"
)

// updateClientCmd represents the updateClient command
var updateClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Update a client's config to that of a JSON file represenign a Keycloak 'ClientRepresentation'",
	Long: `The payload should reflect the Client struct as defined here: https://github.com/Nerzal/gocloak/blob/fe4f627eaf1bff988ee5df2fd0d0b87daac6c074/models.go#L435

Example: keycloak-commander update client --json /path/to/file.json`,
	Run: func(cmd *cobra.Command, args []string) {
		jsonFilename := cmd.Flag("json").Value.String()
		filePayload, err := os.ReadFile(jsonFilename)
		if err != nil {
			panic(err)
		}
		clientRepresentation := &gocloak.Client{}

		err = json.Unmarshal(filePayload, clientRepresentation)
		if err != nil {
			panic(err)
		}
		KeycloakCommander.UpdateClient(args[0], clientRepresentation)
	},
}

func init() {
	updateCmd.AddCommand(updateClientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateClientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	updateClientCmd.Flags().String("json", "", "Json file representing the ClientRepresentation payload.")
	updateClientCmd.MarkFlagRequired("json")
}
