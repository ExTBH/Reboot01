package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"sandbox/internal/server"
)

func main() {
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		portEnv = "8080"
	}
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		log.Printf("Error getting port: %s\n", err.Error())
		os.Exit(1)
	}
	// Ports can only be > 0 && <= 65535
	if !(port > 0 && port <= math.MaxUint16) {
		log.Printf("Invalid port: %s\n", portEnv)
		os.Exit(1)
	}
	server.GoLive(portEnv)
}
