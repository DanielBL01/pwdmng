package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const db = "password.db"

// 'add' functionality
func add(website string, username string, password string) {
	entry := website + "," + username + "," + password + "\n"

	/* OpenFile is Open but with flags
	O_RDONLY	It opens the file read-only
	O_WRONLY	It opens the file write-only
	O_RDWR		It opens the file read-write
	O_APPEND	It appends data to the file when writing
	O_CREATE	It creates a new file if none exists
	O_EXCL		It is used with O_CREATE, and the file must not exist
	O_SYNC		It opens for synchronous I/O
	O_TRUNC		if possible, truncate the file when opened */
	f, err := os.OpenFile(db, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.WriteString(entry); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	statement := "Added " + website + " credentials to pwdmng"
	fmt.Println(statement)
}

// 'get' functionality
func get(website string) {
	f, err := os.Open(db)
	if err != nil {
		log.Fatal(err)
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		entry := strings.Split(input.Text(), ",")
		if entry[0] == website {
			fmt.Println(entry[1], entry[2])
			return
		}
	}

	statement := "Could not find " + website + " credentials in pwdmng"
	fmt.Println(statement)
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
	default:
		fmt.Println("Invalid Operation")
	}
}
