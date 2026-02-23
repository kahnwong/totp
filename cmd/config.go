package cmd

import (
	"log"

	cliBase "github.com/kahnwong/cli-base-sops"
)

var config *Config

func init() {
	var err error
	config, err = cliBase.ReadYamlSops[Config]("~/.config/totp/totp.sops.yaml")
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
}

type Config struct {
	Totp []struct {
		Org      string `yaml:"org"`
		Accounts []struct {
			Name  string `yaml:"name"`
			Token string `yaml:"token"`
		} `yaml:"accounts"`
	} `yaml:"totp"`
}
