package main

import (
	"groupie-tracker/internal/geocoderapi"
	"groupie-tracker/internal/webapp"
	"io"
	"log"
	"os"
)

func main() {
	setupLogger()
	geocoderapi.SetMapBoxToken("YOUR-MAPBOX-TOKEN")
	err := webapp.ListenOn(8080)
	if err != nil {
		log.Fatalln(err)
	}
}

func setupLogger() {
	// Set the TZ environment variable to "UTC" to use UTC (UTC+0) time zone.
	os.Setenv("TZ", "UTC")
	defer os.Setenv("TZ", "") // Reset the TZ environment variable when done.

	// Open the log file for writing. This will create a new file or truncate
	// the existing one if it already exists.
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}

	// Create a multi-writer that writes log messages both to the file and the terminal.
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Set up the logger to use the multi-writer and include timestamps.
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime)
}
