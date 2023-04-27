package main

import _ "modernc.org/sqlite"
import "database/sql"
import "fmt"
import "log"

func main() {

	db, err := sql.Open("sqlite", ":memory:")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var version string
	err = db.QueryRow("EXPLAIN SELECT * FROM libsql_wasm_func_table").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}
