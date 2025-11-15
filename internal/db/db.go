package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Conn *sql.DB
}

func New(path string) (*Database, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(1)

	// Test connection
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	db := &Database{Conn: conn}

	// Run migrations
	if err := db.migrate(); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Database) migrate() error {
	// Create health_check table
	query := `
		CREATE TABLE IF NOT EXISTS health_check (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			checked_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err := db.Conn.Exec(query)
	return err
}

func (db *Database) Close() error {
	return db.Conn.Close()
}
