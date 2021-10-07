package rdb

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Local Redis runs on 127.0.0.1:6379
var db = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func Store(website string, username string, password string) {
	err := db.HSet(ctx, website, []string{"username", username, "password", password}).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully added: " + website)
}

func Retrieve(website string) {
	username, err1 := db.HGet(ctx, website, "username").Result()
	password, err2 := db.HGet(ctx, website, "password").Result()
	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("username: " + username + "\n" + "password: " + password)
}

// Should have an additional confirmation
func Delete(website string) {
	err := db.Del(ctx, website).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully deleted: " + website)
}

// Delete entire database - requires additional confirmation
func Terminate() {
	err := db.FlushAll(ctx).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully terminated all data from pwdmng")
}
