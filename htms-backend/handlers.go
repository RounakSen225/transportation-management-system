package main

import (
	"database/sql"
	"fmt"
)

// Fetch and print all clients from the database
func fetchClients(db *sql.DB) ([]Client, error) {
	query := `SELECT id, name, address, budget FROM clients;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying clients: %w", err)
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var c Client
		err := rows.Scan(&c.ID, &c.Name, &c.Address, &c.Budget)
		if err != nil {
			return nil, fmt.Errorf("error scanning client: %w", err)
		}
		clients = append(clients, c)
	}
	return clients, nil
}

// Fetch and print all transportation events from the database
func fetchTransportationEvents(db *sql.DB) ([]TransportationEvent, error) {
	query := `SELECT id, client_id, date, cost, service_provider FROM transportation_events;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying transportation events: %w", err)
	}
	defer rows.Close()

	var events []TransportationEvent
	for rows.Next() {
		var e TransportationEvent
		err := rows.Scan(&e.ClientID, &e.Date, &e.Cost, &e.ServiceProvider)
		if err != nil {
			return nil, fmt.Errorf("error scanning transportation event: %w", err)
		}
		events = append(events, e)
	}
	return events, nil
}

/*Uncomment to view data in database
func main() {
	db := dbConnect()
	defer db.Close()

	// Fetch and print all clients
	clients, err := fetchClients(db)
	if err != nil {
		log.Fatalf("Failed to fetch clients: %v", err)
	}
	for _, client := range clients {
		fmt.Printf("Client ID: %d, Name: %s, Address: %s, Budget: %.2f\n", client.ID, client.Name, client.Address, client.Budget)
	}

	// Fetch and print all transportation events
	events, err := fetchTransportationEvents(db)
	if err != nil {
		log.Fatalf("Failed to fetch transportation events: %v", err)
	}
	for _, event := range events {
		fmt.Printf("Client ID: %d, Date: %v, Cost: %.2f, Provider: %s\n", event.ClientID, event.Date, event.Cost, event.ServiceProvider)
	}
}
*/
