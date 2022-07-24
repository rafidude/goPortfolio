package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/testdb")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	sql := "INSERT INTO cities(name, population) VALUES ('Moscow', 12506000)"
	row, err := db.Exec(sql)
	if err != nil {
		panic(err.Error())
	}

	lastId, err := row.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
}
