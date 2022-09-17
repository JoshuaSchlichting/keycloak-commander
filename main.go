/*
Copyright © 2022 Joshua Schlichting 7680545+JoshuaSchlichting@users.noreply.github.com

*/
package main

import (
	"errors"
	"keycloak-commander/cmd"
	"log"
	"os"
)

func main() {
	cmd.ConfigFileWriter = configFileWriter
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
