/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/kellyma2/go-boc-valet/pkg"
	"github.com/spf13/cobra"
)

// groupsCmd represents the groups command
var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "Subcommands related to the groups endpoint",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		c := pkg.NewValetClient()
		r, err := c.Group(name)
		if err != nil {
			fmt.Println(err)
		} else {
			r.PrettyPrint()
		}
	},
}

func init() {
	rootCmd.AddCommand(groupsCmd)
}
