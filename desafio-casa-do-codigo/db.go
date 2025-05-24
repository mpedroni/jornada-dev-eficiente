package main

import (
	"database/sql"
	"os"
	"path/filepath"
)

func openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:tmp/db.sqlite")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	schema, err := os.ReadFile(filepath.Join(wd, "db", "schema.sql"))
	if err != nil {
		return err
	}

	if _, err := db.Exec(string(schema)); err != nil {
		return err
	}

	return nil
}
