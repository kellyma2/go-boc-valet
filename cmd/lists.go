package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listsCmd = &cobra.Command{
	Use:   "lists",
	Short: "Subcommands related to lists endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lists called")
	},
}

func init() {
	rootCmd.AddCommand(listsCmd)
}
