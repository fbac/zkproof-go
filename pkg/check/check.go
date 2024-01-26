package check

import (
	"strconv"
	"strings"
)

// Check valid string
func IsValidString(s string) bool {
	return len(strings.TrimSpace(s)) > 0
}

// Port basic sanity checks
func IsValidPort(s string) bool {
	p, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	if p < 1024 {
		return false
	}

	return true
}
