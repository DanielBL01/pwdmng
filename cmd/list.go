package cmd

import (
	"pwdmng/rdb"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all keys",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func list() {
	rdb.List()
}
