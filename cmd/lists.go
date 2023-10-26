package cmd

import (
	"github.com/spf13/cobra"
)

var listsCmd = &cobra.Command{
	Use:   "lists",
	Short: "Subcommands related to lists endpoint",
}

func init() {
	rootCmd.AddCommand(listsCmd)
}
