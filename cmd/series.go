package cmd

import (
	"fmt"

	"github.com/kellyma2/go-boc-valet/pkg"
	"github.com/spf13/cobra"
)

var seriesCmd = &cobra.Command{
	Use:   "series",
	Short: "Subcommands related to lists/series endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		c := pkg.NewValetClient()
		r, err := c.SeriesList()
		if err != nil {
			fmt.Println(err)
		} else {
			for _, info := range r {
				info.PrettyPrint()
			}
		}
	},
}

func init() {
	listsCmd.AddCommand(seriesCmd)
}
