package main

import (
	"fmt"
	"os"
	"pwdmng/rdb"
)

// 'add' functionality
func add(website string, username string, password string) {
	rdb.Store(website, username, password)
}

// 'get' functionality
func get(website string) {
	rdb.Retrieve(website)
}

// 'delete' functionality
func delete(website string) {
	rdb.Delete(website)
}

// 'terminate' functionality
func terminate() {
	rdb.Terminate()
}

func main() {
	switch args := os.Args; args[1] {
	case "add":
		if len(args) == 5 {
			add(args[2], args[3], args[4])
		} else {
			fmt.Println("Invalid arguments")
		}
	case "get":
		if len(args) == 3 {
			get(args[2])
		} else {
			fmt.Println("Invalid arguments")
		}
	case "del":
		if len(args) == 3 {
			delete(args[2])
		}
	case "terminate":
		if len(args) == 2 {
			terminate()
		}
	default:
		fmt.Println("Invalid Operation")
	}
}
