package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init() error {
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}

	_, err = db.Exec(`
                CREATE TABLE IF NOT EXISTS clients (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        name TEXT,
                        email TEXT
                )
        `)
	if err != nil {
		return err
	}
	return nil
}

func InsertClient(name, email string) error {
	_, err := db.Exec("INSERT INTO clients (name, email) VALUES (?, ?)", name, email)
	if err != nil {
		return err
	}
	return nil
}

func QueryClients() ([]Client, error) {
	rows, err := db.Query("SELECT * FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var client Client
		err := rows.Scan(&client.ID, &client.Name, &client.Email)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func UpdateClient(id int, newEmail string) error {
	_, err := db.Exec("UPDATE clients SET email = ? WHERE id = ?", newEmail, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteClient(id int) error {
	_, err := db.Exec("DELETE FROM clients WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

type Client struct {
	ID    int
	Name  string
	Email string
}
