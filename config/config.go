package config

import (
	"os"
	"strconv"
)

const (
	// Version of the application
	Version = "2.7"

	// DefaultPort is the default port the server will listen on
	DefaultPort = "80"
)

// GetPort returns the port the server should listen on
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":" + DefaultPort
	}

	// If port is already in format ":8080", return as is
	if port[0] == ':' {
		return port
	}

	// Check if port is a valid number
	if _, err := strconv.Atoi(port); err != nil {
		return ":" + DefaultPort
	}

	return ":" + port
}
