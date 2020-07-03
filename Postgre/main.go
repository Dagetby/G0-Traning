package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	connStr := "user=postgres password=example dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	check(err)

	defer db.Close()

	result, err := db.Exec("INSERT INTO users (first_name, last_name) VALUES ($1, $2)", "Nikolas", "Keigh")

	check(err)

	fmt.Println(result.LastInsertId())

}
