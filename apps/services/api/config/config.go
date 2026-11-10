// Package config reads the service's settings from the environment.
package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds everything the service reads from the environment.
type Config struct {
	Port           string
	AllowedOrigins []string
}

// Load reads the environment once and reports the first missing or invalid variable.
func Load() (Config, error) {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = "8080"
	}
	if n, err := strconv.Atoi(port); err != nil || n < 1 || n > 65535 {
		return Config{}, fmt.Errorf("PORT must be a number from 1 to 65535, got %q", port)
	}

	var origins []string
	for _, o := range strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",") {
		if o = strings.TrimSpace(o); o != "" {
			origins = append(origins, o)
		}
	}
	if len(origins) == 0 {
		return Config{}, errors.New("ALLOWED_ORIGINS is required")
	}

	return Config{Port: port, AllowedOrigins: origins}, nil
}
