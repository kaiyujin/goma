package goma

import (
	"database/sql"
	"log"
)

func Query(db *sql.DB, path string) {
	sql, _ := getFile(path)
	stmt := prepareStmt(db, sql)
	stmt.Query()
}

func prepareStmt(db *sql.DB, query string) *sql.Stmt {
	if db == nil {
		return nil
	}
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal("Could not prepare `" + query + "`: " + err.Error())
	}
	return stmt
}
