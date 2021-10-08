package cmd

import (
	"fmt"
	"pwdmng/rdb"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get your credentials for a website/platform",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			get(args[0])
		} else {
			fmt.Println("Invalid arguments")
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}

func get(website string) {
	rdb.Retrieve(website)
}
