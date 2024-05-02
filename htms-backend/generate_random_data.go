package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

// Helper functions for random data generation
func randomName() string {
	firstNames := []string{"John", "Jane", "Alice", "Bob", "Carol", "David"}
	lastNames := []string{"Smith", "Doe", "Johnson", "White", "Brown", "Davis"}
	return firstNames[rand.Intn(len(firstNames))] + " " + lastNames[rand.Intn(len(lastNames))]
}

func randomAddress() string {
	streets := []string{"Maple Street", "Oak Street", "Pine Street", "Elm Street", "Cedar Street"}
	numbers := rand.Intn(1000)
	return fmt.Sprintf("%d %s", numbers, streets[rand.Intn(len(streets))])
}

func randomBudget() float64 {
	return float64(1000 + rand.Intn(9000))
}

func randomTransportationEvent(clientID int) TransportationEvent {
	services := []string{"Uber", "Lyft", "Local Taxi", "Health Transport"}
	cost := float64(10 + rand.Intn(90))
	daysAgo := rand.Intn(30) // events up to 30 days ago
	date := time.Now().AddDate(0, 0, -daysAgo)

	return TransportationEvent{
		ClientID:        clientID,
		Date:            date,
		Cost:            cost,
		ServiceProvider: services[rand.Intn(len(services))],
	}
}

func insertSampleData(db *sql.DB) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		name := randomName()
		address := randomAddress()
		budget := randomBudget()

		// Using QueryRow to execute the insert statement and capture the last inserted ID
		var clientId int
		err := db.QueryRow("INSERT INTO clients (name, address, budget) VALUES ($1, $2, $3) RETURNING id",
			name, address, budget).Scan(&clientId)
		if err != nil {
			fmt.Printf("Error inserting client: %v\n", err)
			continue
		}

		// Each client will have between 5 and 10 transportation events
		eventCount := 5 + rand.Intn(6)
		for j := 0; j < eventCount; j++ {
			event := randomTransportationEvent(clientId)
			_, err := db.Exec("INSERT INTO transportation_events (client_id, date, cost, service_provider) VALUES ($1, $2, $3, $4)",
				clientId, event.Date, event.Cost, event.ServiceProvider)
			if err != nil {
				fmt.Printf("Error inserting transportation event for client %d: %v\n", clientId, err)
				continue
			}
		}
	}

	fmt.Println("Sample data insertion complete.")
}
