package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
	"golang.org/x/exp/slices"
)

func TOTPGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var autocompleteOptions []string

	if len(args) == 0 { // organizations
		autocompleteOptions = getOrgs()
	} else if len(args) == 1 { // accounts
		autocompleteOptions = getAccounts(args[0])
	}

	return autocompleteOptions, cobra.ShellCompDirectiveNoFileComp
}

var generateCmd = &cobra.Command{
	Use:               "generate [org] [account]",
	Short:             "Generate TOTP",
	Long:              `Generate TOTP`,
	ValidArgsFunction: TOTPGet,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify an organization and account")
			os.Exit(1)
		} else if len(args) == 1 {
			fmt.Println("Please specify an account")
			os.Exit(1)
		}

		// generate TOTP
		var secret string
		for _, c := range config {
			if c.Org == args[0] {
				for _, account := range c.Accounts {
					if account.Name == args[1] {
						secret = account.Token
					}
					continue
				}
				continue
			}
		}
		isValidOrg := slices.Contains(getOrgs(), args[0]) // true

		if isValidOrg {
			isValidAccount := slices.Contains(getAccounts(args[0]), args[1]) // true
			if isValidAccount {
				passcode, err := totp.GenerateCode(secret, time.Now())
				if err != nil {
					panic(err)
				}

				fmt.Println(passcode)

				// copy to clipboard
				err = clipboard.Init()
				if err != nil {
					panic(err)
				}
				clipboard.Write(clipboard.FmtText, []byte(passcode))
			} else {
				fmt.Println("Please specify an available account")
				os.Exit(1)
			}
		} else {
			fmt.Println("Please specify an available organization")
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

type ConnectCommand struct{}

func (c *ConnectCommand) Help() string {
	return "Generate TOTP"
}

func (c *ConnectCommand) Synopsis() string {
	return "Generate TOTP"
}
