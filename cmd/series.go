package cmd

import (
	"fmt"

	"github.com/kellyma2/go-boc-valet/pkg"
	"github.com/spf13/cobra"
)

var seriesCmd = &cobra.Command{
	Use:   "series <name>",
	Short: "Subcommands related to the series endpoint",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		c := pkg.NewValetClient()
		r, err := c.Series(name)
		if err != nil {
			fmt.Println(err)
		} else {
			r.PrettyPrint()
		}
	},
}

func init() {
	rootCmd.AddCommand(seriesCmd)
}
