package cmd

import (
	"fmt"

	"github.com/kellyma2/go-boc-valet/pkg"
	"github.com/spf13/cobra"
)

var observationsSeriesCmd = &cobra.Command{
	Use:   "series",
	Short: "Operations related to the series observations endpoint",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		c := pkg.NewValetClient()
		r, err := c.SeriesObservations(pkg.WithSeries(name), pkg.WithRecent(10))
		if err != nil {
			fmt.Println(err)
		} else {
			r.PrettyPrint()
		}
	},
}

func init() {
	observationsCmd.AddCommand(observationsSeriesCmd)
}
