// Package config reads the API's settings from environment variables.
package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds everything the API reads from its environment.
type Config struct {
	Port           string
	AllowedOrigins []string
}

// Load reads the environment once; a missing or invalid variable is an error naming it.
func Load() (Config, error) {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = "8080"
	}
	if n, err := strconv.Atoi(port); err != nil || n < 1 || n > 65535 {
		return Config{}, fmt.Errorf("invalid PORT %q: must be a number from 1 to 65535", port)
	}

	var origins []string
	for _, origin := range strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",") {
		if origin = strings.TrimSpace(origin); origin != "" {
			origins = append(origins, origin)
		}
	}
	if len(origins) == 0 {
		return Config{}, errors.New("missing required environment variable ALLOWED_ORIGINS")
	}

	return Config{Port: port, AllowedOrigins: origins}, nil
}
