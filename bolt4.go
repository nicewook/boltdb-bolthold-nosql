package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/gommon/log"
	"github.com/timshannon/bolthold"
)

type User struct {
	ID int `boltholdKey:"ID"`
	// Group string
	Email string
	Name  string
	Age   int
}

const databaseFile = "key-value.db"

func main() {
	// https://github.com/timshannon/bolthold

	// 1. open
	store, err := bolthold.Open(databaseFile, 0666, nil)
	if err != nil {
		log.Fatalf("fail to open db: %v", err)
	}
	defer store.Close()

	// 2. fetch all objects
	var users []User
	if err := store.Find(&users, nil); err != nil {
		fmt.Println(err)
	}

	// 5. print
	spew.Dump(users)
}

/*
([]main.User) (len=3 cap=4) {
	(main.User) {
		ID: (int) 1,
		Email: (string) "",
		Name: (string) (len=8) "John Doe",
		Age: (int) 21
	},
	(main.User) {
		ID: (int) 2,
		Email: (string) (len=13) "aaa@gmail.com",
		Name: (string) (len=8) "John Doe",
		Age: (int) 21
	},
	(main.User) {
		ID: (int) 3,
		Email: (string) (len=13) "bbb@gmail.com",
		Name: (string) "",
		Age: (int) 33
	}
}
*/
