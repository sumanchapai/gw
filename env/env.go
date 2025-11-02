package env

import (
	"os"
	"strings"
)

// Host returns the value of the Host environment variable, and default if not set
func Host() string {
	if v := os.Getenv("Host"); v != "" {
		return v
	}
	return "0.0.0.0"
}

// PORT returns the value of the Port environment variable, and default if not set
func Port() string {
	if v := os.Getenv("Port"); v != "" {
		return v
	}
	return "8000"
}

// GW_REPO return the value of the GW_REPO environment variable
func GW_REPO() string {
	return os.Getenv("GW_REPO")
}

// BACK_LINK returns the value of the Back_Link environment variable
func BACK_LINK() string {
	return os.Getenv("BACK_LINK")
}

// BACK_LINK returns the value of the Back_Link environment variable
func BASE_PATH() string {
	return os.Getenv("BASE_PATH")
}

// RESTRICTED_COMMANDS returns the list of restricted git commands
func RESTRICTED_COMMANDS() []string {
	cmds := os.Getenv("RESTRICTED_COMMANDS")
	return strings.Split(cmds, ",")
}

// Title returns the value of GIT_PAGE_TITLE environment variable or a default value if not set
func Title() string {
	if title := os.Getenv("GIT_PAGE_TITLE"); title != "" {
		return title
	}
	return "Git Web"
}
