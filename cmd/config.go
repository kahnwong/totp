package cmd

import (
	cliBase "github.com/kahnwong/cli-base-sops"
)

var config = cliBase.ReadYamlSops[Config]("~/.config/totp/totp.sops.yaml")

type Config struct {
	Totp []struct {
		Org      string `yaml:"org"`
		Accounts []struct {
			Name  string `yaml:"name"`
			Token string `yaml:"token"`
		} `yaml:"accounts"`
	} `yaml:"totp"`
}
