package db

import (
	"database/sql"
	"fmt"

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
	// query := `
	// 	CREATE TABLE IF NOT EXISTS health_check (
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		checked_at DATETIME DEFAULT CURRENT_TIMESTAMP
	// 	);
	// `
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS oauth_clients (
	  client_name TEXT NOT NULL UNIQUE,
      client_id TEXT NOT NULL UNIQUE PRIMARY KEY,
      client_secret TEXT NOT NULL,
      name TEXT NOT NULL,
      grant_types TEXT NOT NULL,
      scopes TEXT,
      is_confidential INTEGER DEFAULT 1,
      redirect_uris TEXT,
      metadata TEXT,
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
      updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
  );
  `,
		`
		CREATE TABLE IF NOT EXISTS health_check (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			checked_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`,
	}

	for i, sql := range migrations {
		if _, err := db.Conn.Exec(sql); err != nil {
			return fmt.Errorf("migration %d failed: %w", i+1, err)
		}
	}

	return nil

}

func (db *Database) Close() error {
	return db.Conn.Close()
}
