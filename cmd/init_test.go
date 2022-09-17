package cmd

import (
	"testing"
)

func TestInitOutputContainsFlags(t *testing.T) {
	initCmd.Flag("host").Value.Set("http://localhost:8080/")
	initCmd.Flag("realm").Value.Set("master")
	initCmd.Flag("username").Value.Set("adminname")
	initCmd.Flag("password").Value.Set("adminpass")

	expectedPayload := []byte(`{"host":"http://localhost:8080/","realm":"master","username":"adminname","password":"adminpass"}`)
	configOutput := []byte{}
	ConfigFileWriter = func(data []byte) error {
		configOutput = data
		return nil
	}

	initCmd.Run(initCmd, []string{})
	if string(configOutput) != string(expectedPayload) {
		t.Errorf("CommanderConfig is not receving the expected payload, %s", string(expectedPayload))
	}
}
