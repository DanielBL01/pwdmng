package cmd

import (
	"pwdmng/rdb"

	"github.com/spf13/cobra"
)

var terminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: "Erase all data in pwdmng",
	Run: func(cmd *cobra.Command, args []string) {
		terminate()
	},
}

func init() {
	RootCmd.AddCommand(terminateCmd)
}

func terminate() {
	rdb.Terminate()
}
