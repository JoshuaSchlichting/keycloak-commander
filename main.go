package main

import (
	"encoding/json"
	"errors"
	"keycloak-commander/cmd"
	"keycloak-commander/keycloak"
	"log"
	"os"
)

func main() {
	cmd.ConfigFileWriter = configFileWriter
	config := getConfigFromFile(getConfigFilename())
	// TODO: Don't require the commander until a command is actually run
	cmd.KeycloakCommander = keycloak.NewKeycloakCommander(
		config.Host,
		config.Username,
		config.Password,
		config.Realm,
	)
	cmd.Execute()
}

func getConfigFilename() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return home + "/.keycloak-commander.json"
}

func createConfigFile() (filename string) {
	// create config file in home directory if the file doesn't exist
	if !getConfigFileExists() {
		_, err := os.Create(getConfigFilename())
		if err != nil {
			log.Fatal(err)
		}
	}
	return getConfigFilename()
}

func getConfigFileExists() bool {
	if _, err := os.Stat(getConfigFilename()); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		panic(err)
	}
}

func configFileWriter(data []byte) error {
	// write data to file
	file, err := os.OpenFile(getConfigFilename(), os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFromFile(filename string) cmd.Config {
	// read config file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	config := cmd.Config{}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
