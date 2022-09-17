package cmd

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
)

var ConfigFileWriter func(data []byte) error

type Config struct {
	Host     string `json:"host"`
	Realm    string `json:"realm"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Creating config file")
		host := cmd.Flag("host").Value.String()
		realm := cmd.Flag("realm").Value.String()
		username := cmd.Flag("username").Value.String()
		password := cmd.Flag("password").Value.String()
		if username == "" || password == "" {
			log.Fatal("Username and password are required")
		}
		config := Config{
			Host:     host,
			Realm:    realm,
			Username: username,
			Password: password,
		}
		configPayload, err := json.Marshal(config)
		if err != nil {
			log.Fatal(err)

		}
		ConfigFileWriter(configPayload)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().StringP("host", "h", "", "Hostname of the Keycloak server")
	initCmd.Flags().StringP("realm", "r", "", "Realm to use")
	initCmd.Flags().StringP("username", "u", "", "Admin's username")
	initCmd.Flags().StringP("password", "p", "", "Admin's password")
}
