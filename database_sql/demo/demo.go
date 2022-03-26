package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./somnus.db")
	if err != nil {
		fmt.Println(err)
		log.Panicln("db found failed")
	}
	defer func() { _ = db.Close() }()
	_, err = db.Exec("DROP TABLE IF EXISTS User;")
	if err != nil {
		log.Panicln("Drop table failed")
	}
	_, err = db.Exec("CREATE TABLE User(Name text);")
	if err != nil {
		log.Panicln("Create table failed")
	}
	_, err = db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
	if err != nil {
		log.Panicln("Insert table failed")
	}
	var name string
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}
