package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of organizations",
	Long:  `Get a list of organizations`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a command: [orgs | accounts]")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
