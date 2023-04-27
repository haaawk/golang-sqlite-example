package main

import _ "github.com/mattn/go-sqlite3"
import "database/sql"
import "fmt"
import "log"

func main() {

	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}
