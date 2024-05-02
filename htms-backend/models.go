package main

import (
	"time"
)

type Client struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Budget  float64 `json:"budget"`
}

type TransportationEvent struct {
	ClientID        int       `json:"clientID"`
	Date            time.Time `json:"date"`
	Cost            float64   `json:"cost"`
	ServiceProvider string    `json:"serviceProvider"`
}

// Define the Alert struct
type Alert struct {
	ID        int     `json:"id"`
	ClientID  int     `json:"clientID"`
	Message   string  `json:"message"`
	Triggered bool    `json:"triggered"`
	Value     float64 `json:"value"`
}
