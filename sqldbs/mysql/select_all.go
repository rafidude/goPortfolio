package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	Id         int
	Name       string
	Population int
}

func main() {

	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/testdb")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM cities")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var city City
		err := rows.Scan(&city.Id, &city.Name, &city.Population)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", city)
	}
}
