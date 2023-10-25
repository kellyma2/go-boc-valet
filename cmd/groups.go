package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Subcommands related to lists/groups endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("groups called")
	},
}

func init() {
	listsCmd.AddCommand(groupsCmd)
}
