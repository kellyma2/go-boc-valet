package cmd

import (
	"fmt"

	"github.com/kellyma2/go-boc-valet/pkg"
	"github.com/spf13/cobra"
)

var observationsGroupCmd = &cobra.Command{
	Use:   "group",
	Short: "Operations related to the group observations endpoint",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		c := pkg.NewValetClient()
		r, err := c.GroupObservations(pkg.WithItem(name), pkg.WithRecent(5))
		if err != nil {
			fmt.Println(err)
		} else {
			r.PrettyPrint()
		}
	},
}

func init() {
	observationsCmd.AddCommand(observationsGroupCmd)
}
