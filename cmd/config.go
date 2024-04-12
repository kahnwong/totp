package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/getsops/sops/v3/decrypt"

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
	filename := filepath.Join(homeDir, ".config", "totp", "totp.sops.yaml.txt")

	// Check if the file exists
	_, err = os.Stat(filename)

	if os.IsNotExist(err) {
		fmt.Printf("File %s does not exist.\n", filename)
		os.Exit(1)
	}

	var configs []Config

	data, err := decrypt.File(filename, "txt") // sops yaml specs does not support array at root
	if err != nil {
		fmt.Println(fmt.Printf("Failed to decrypt: %v", err))
	}

	err = yaml.Unmarshal(data, &configs)
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

func getAccounts(org string) []string {
	accounts := make([]string, 0)
	for _, v := range config {
		if v.Org == org {
			for _, v := range v.Accounts {
				accounts = append(accounts, v.Name)
			}
		}
	}

	return accounts
}
