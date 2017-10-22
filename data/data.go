package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
	fmt.Println(Db, err)
	if err != nil {
		log.Fatal(err)
	}

	return
}
