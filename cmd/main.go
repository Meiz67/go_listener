package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	fmt.Println("Listener started")
	db, _ := sql.Open("postgres", "port=5432 host=postgres user=postgres password=root dbname=postgres sslmode=disable")
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected")
	}
}
