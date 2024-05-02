package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Function to check for budget alerts for a specific client
func checkAndTriggerAlerts(db *sql.DB, clientID string) ([]Alert, error) {
	var alerts []Alert

	// Placeholder for logic to check client's spending against budget
	// This should be replaced with your actual logic
	var clientBudget, clientSpending float64
	var clientIDInt int
	err := db.QueryRow("SELECT budget FROM clients WHERE id = $1", clientID).Scan(&clientBudget)
	if err != nil {
		return nil, err
	}
	err = db.QueryRow("SELECT SUM(cost) FROM transportation_events WHERE client_id = $1", clientID).Scan(&clientSpending)
	if err != nil {
		return nil, err
	}
	clientIDInt, _ = strconv.Atoi(clientID)
	threshold := 0.8 // Example threshold for alert
	if clientSpending/clientBudget >= threshold {
		alert := Alert{
			ClientID:  clientIDInt, // Ensure this conversion is safe
			Message:   "Budget threshold reached",
			Triggered: true,
			Value:     clientSpending,
		}
		alerts = append(alerts, alert)
	}

	return alerts, nil
}

func main() {
	db := dbConnect()
	defer db.Close()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow only the frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Preflight requests can be cached for 12 hours
	}))

	router.GET("/clients", func(c *gin.Context) {
		clients, err := fetchClients(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, clients)
	})

	router.GET("/events", func(c *gin.Context) {
		events, err := fetchTransportationEvents(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, events)
	})

	// Get a specific client profile
	router.GET("/clients/:id", func(c *gin.Context) {
		var client Client
		id := c.Param("id")

		// Use a standard SQL query instead of db.Where
		query := `SELECT id, name, address, budget FROM clients WHERE id = $1;`
		row := db.QueryRow(query, id)
		err := row.Scan(&client.ID, &client.Name, &client.Address, &client.Budget)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
		c.JSON(http.StatusOK, client)
	})

	// Record a new transportation event
	router.POST("/events", func(c *gin.Context) {
		var event TransportationEvent
		if err := c.ShouldBindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert the event into the database
		query := `INSERT INTO transportation_events (client_id, date, cost, service_provider) VALUES ($1, $2, $3, $4);`
		_, err := db.Exec(query, event.ClientID, time.Now(), event.Cost, event.ServiceProvider)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, event)
	})

	// Endpoint to check and trigger alerts for a client
	router.GET("/clients/:id/alerts", func(c *gin.Context) {
		clientID := c.Param("id")
		alerts, err := checkAndTriggerAlerts(db, clientID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for alerts"})
			return
		}
		c.JSON(http.StatusOK, alerts)
	})

	router.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
