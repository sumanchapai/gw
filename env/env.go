package env

import (
	"os"
	"strings"
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

// Return the value of the Back_Link environment variable
func BACK_LINK() string {
	return os.Getenv("BACK_LINK")
}

// Return the list of Restricted git commands by RESTRICTED_COMMANDS envvar.
func RESTRICTED_COMMANDS() []string {
	cmds := os.Getenv("RESTRICTED_COMMANDS")
	return strings.Split(cmds, ",")
}
