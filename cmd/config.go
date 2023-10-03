package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// init
var config = readConfig()

// readConfig
type Config struct {
	Org      string `yaml:"org"`
	Accounts []struct {
		Name  string `yaml:"name"`
		Token string `yaml:"token"`
	} `yaml:"accounts"`
}

func readConfig() []Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	filename := filepath.Join(homeDir, ".config", "totp", "totp.yaml")

	// Check if the file exists
	_, err = os.Stat(filename)

	if os.IsNotExist(err) {
		fmt.Printf("File %s does not exist.\n", filename)
		os.Exit(1)
	}

	var configs []Config

	source, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("failed reading config file: %v\n", err)
	}

	err = yaml.Unmarshal(source, &configs)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return configs
}

func getOrgs() []string {
	orgs := make([]string, 0)
	for _, v := range config {
		orgs = append(orgs, v.Org)
	}

	return orgs
}
