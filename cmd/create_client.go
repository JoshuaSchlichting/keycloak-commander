package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Create client(s).",
	Long: `Each argument shall be a string represent a client name. At least one client name is required.
	
Example: keycloak-commander create client my-client another-client a-third-client`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("At least one client name is required.")
		}
		for _, arg := range args {
			log.Printf("Creating client %s\n", arg)
			KeycloakCommander.CreateClient(arg)
		}
	}}

func init() {
	createCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
