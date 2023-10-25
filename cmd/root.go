package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-boc-valet",
	Short: "Bank of Canada Valet API wrapper",
	Long: `go-boc-valet provides both a library and a trivial CLI
wrapper that allows access to the Bank of Canada Valet API.

See: https://www.bankofcanada.ca/valet/docs#valet_api`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
