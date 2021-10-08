package cmd

import (
	"fmt"
	"pwdmng/rdb"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Store your credentials for a website/platform",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 3 {
			add(args[0], args[1], args[2])
		} else {
			fmt.Println("Invalid arguments")
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

func add(website string, username string, password string) {
	rdb.Store(website, username, password)
}
