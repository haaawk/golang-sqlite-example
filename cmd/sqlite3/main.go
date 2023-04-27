package main

import _ "github.com/psarna/go-sqlite3"
import "database/sql"
import "fmt"
import "log"

func main() {

	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

    var addr sql.NullInt64
	var opcode sql.NullString
    var p1, p2, p3, p4, p5 sql.NullInt64
    var comment sql.NullString
	err = db.QueryRow("EXPLAIN SELECT * FROM libsql_wasm_func_table").Scan(&addr, &opcode, &p1, &p2, &p3, &p4, &p5, &comment)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addr, opcode, p1, p2, p3, p4, p5, comment)
}
