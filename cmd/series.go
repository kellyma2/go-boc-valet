package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var seriesCmd = &cobra.Command{
	Use:   "series",
	Short: "Subcommands related to lists/series endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("series called")
	},
}

func init() {
	listsCmd.AddCommand(seriesCmd)
}
