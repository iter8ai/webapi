// Package config reads the API's settings from environment variables.
package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		return Config{}, fmt.Errorf("PORT must be a number from 1 to 65535, got %q", port)
	}

	var origins []string
	for _, o := range strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",") {
		if o = strings.TrimSpace(o); o != "" {
			origins = append(origins, o)
		}
	}
	if len(origins) == 0 {
		return Config{}, errors.New("ALLOWED_ORIGINS is required: a comma-separated list of origins")
	}

	return Config{Port: port, AllowedOrigins: origins}, nil
}
