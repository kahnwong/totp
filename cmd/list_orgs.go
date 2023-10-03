package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var orgsCmd = &cobra.Command{
	Use:   "orgs",
	Short: "Get a list of organizations",
	Long:  `Get a list of organizations`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("Available orgs:")
		for _, v := range getOrgs() {
			fmt.Printf("  - %s\n", v)
		}
	},
}

func init() {
	listCmd.AddCommand(orgsCmd)
}
