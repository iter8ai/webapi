package config

import (
	"strings"
	"testing"
)

func TestLoadMissingAllowedOrigins(t *testing.T) {
	for _, value := range []string{"", "  "} {
		t.Setenv("ALLOWED_ORIGINS", value)
		_, err := Load()
		if err == nil || !strings.Contains(err.Error(), "ALLOWED_ORIGINS") {
			t.Fatalf("value %q: want error naming ALLOWED_ORIGINS, got %v", value, err)
		}
	}
}

func TestLoadInvalidPort(t *testing.T) {
	t.Setenv("ALLOWED_ORIGINS", "http://localhost:5173")
	for _, value := range []string{"abc", "0", "70000"} {
		t.Setenv("PORT", value)
		_, err := Load()
		if err == nil || !strings.Contains(err.Error(), "PORT") {
			t.Fatalf("value %q: want error naming PORT, got %v", value, err)
		}
	}
}
