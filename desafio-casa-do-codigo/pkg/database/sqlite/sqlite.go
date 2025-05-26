package sqlite

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *sql.DB, schemaPath string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	schema, err := os.ReadFile(filepath.Join(wd, schemaPath))
	if err != nil {
		return err
	}

	if _, err := db.Exec(string(schema)); err != nil {
		return err
	}

	return nil
}
