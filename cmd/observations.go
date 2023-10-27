package cmd

import (
	"github.com/spf13/cobra"
)

var observationsCmd = &cobra.Command{
	Use:   "observations",
	Short: "Subcommands related to the observations endpoint",
}

func init() {
	rootCmd.AddCommand(observationsCmd)
}
