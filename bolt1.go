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
	// Email string
	Name string
	Age  int
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

	// 2. generate data
	user := User{
		ID:   1,
		Name: "John Doe",
		Age:  21,
	}

	// 3. save. key is "ID" and value is the whole struct
	if err := store.Upsert(user.ID, &user); err != nil {
		log.Fatalf("fail to save record to db: %v", err)
	}

	// 4. fetch all objects
	var users []User
	if err := store.Find(&users, nil); err != nil {
		fmt.Println(err)
	}

	// 5. print
	spew.Dump(users)
}

/*
([]main.User) (len=1 cap=1) {
	(main.User) {
    ID: (int) 1,
    Name: (string) (len=8) "John Doe",
    Age: (int) 21
	}
}
*/
