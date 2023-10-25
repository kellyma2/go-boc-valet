package cmd

import (
	"fmt"

	"github.com/kellyma2/go-boc-valet/pkg"
	"github.com/spf13/cobra"
)

var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Subcommands related to lists/groups endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		c := pkg.NewValetClient()
		r, err := c.GroupList()
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
	listsCmd.AddCommand(groupsCmd)
}
