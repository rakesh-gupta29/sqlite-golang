// database/database.go
package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func Connect(dbPath string) error {
	var err error
	Database, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	_, err = Database.Exec(`
		CREATE TABLE IF NOT EXISTS forms (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			message TEXT NOT NULL
		);
	`)

	if err != nil {
		return err
	}

	return nil
}
