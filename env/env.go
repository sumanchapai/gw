package env

import (
	"os"
)

// Return the value of the Host environment variable, and default if not set
func Host() string {
	if v := os.Getenv("Host"); v != "" {
		return v
	}
	return "0.0.0.0"
}

// Return the value of the Port environment variable, and default if not set
func Port() string {
	if v := os.Getenv("Port"); v != "" {
		return v
	}
	return "8000"
}

// Return the value of the GW_REPO environment variable
func GW_REPO() string {
	return os.Getenv("GW_REPO")
}
