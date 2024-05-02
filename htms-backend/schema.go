package main

import (
	"database/sql"
	"fmt"
	"log"
)

func createTables(db *sql.DB) error {
	clientTable := `CREATE TABLE IF NOT EXISTS clients (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        address VARCHAR(255) NOT NULL,
        budget DECIMAL(10, 2) NOT NULL
    );`
	eventTable := `CREATE TABLE IF NOT EXISTS transportation_events (
        id SERIAL PRIMARY KEY,
        client_id INTEGER NOT NULL,
        date TIMESTAMP NOT NULL,
        cost DECIMAL(10, 2) NOT NULL,
        service_provider VARCHAR(255) NOT NULL,
        FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE
    );`

	_, err := db.Exec(clientTable)
	if err != nil {
		log.Fatalf("Failed to create clients table: %v", err)
	}
	_, err = db.Exec(eventTable)
	if err != nil {
		log.Fatalf("Failed to create transportation_events table: %v", err)
	}
	fmt.Println("Tables created successfully")
	return nil
}
