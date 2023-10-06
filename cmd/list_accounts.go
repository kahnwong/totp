package cmd

import (
	"fmt"
	"os"

	"golang.org/x/exp/slices"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func OrgGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return getOrgs(), cobra.ShellCompDirectiveNoFileComp
}

var accountsCmd = &cobra.Command{
	Use:               "accounts",
	Short:             "Get a list of accounts for a given organization",
	Long:              `Get a list of accounts for a given organization`,
	ValidArgsFunction: OrgGet,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify an organization")
			os.Exit(1)
		}

		isValidOrg := slices.Contains(getOrgs(), args[0]) // true

		if isValidOrg {
			green := color.New(color.FgGreen).SprintFunc()

			fmt.Printf("%s %s\n", green("Organization:"), args[0])
			color.Blue("Accounts:")

			for _, v := range getAccounts(args[0]) {
				fmt.Printf("  - %s\n", v)
			}

		} else {
			fmt.Println("Please specify an available organization")
			os.Exit(1)
		}
	},
}

func init() {
	listCmd.AddCommand(accountsCmd)
}
