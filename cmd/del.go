package cmd

import (
	"fmt"
	"pwdmng/rdb"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete your credentials for a website/platform",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			del(args[0])
		} else {
			fmt.Println("Invalid arguments")
		}
	},
}

func init() {
	RootCmd.AddCommand(delCmd)
}

func del(website string) {
	rdb.Delete(website)
}
