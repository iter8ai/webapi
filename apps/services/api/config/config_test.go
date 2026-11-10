package config

import (
	"strings"
	"testing"
)

func TestLoadMissingAllowedOrigins(t *testing.T) {
	t.Setenv("ALLOWED_ORIGINS", "")
	_, err := Load()
	if err == nil || !strings.Contains(err.Error(), "ALLOWED_ORIGINS") {
		t.Fatalf("expected error naming ALLOWED_ORIGINS, got %v", err)
	}
}

func TestLoadBlankAllowedOrigins(t *testing.T) {
	t.Setenv("ALLOWED_ORIGINS", "  ")
	_, err := Load()
	if err == nil || !strings.Contains(err.Error(), "ALLOWED_ORIGINS") {
		t.Fatalf("expected error naming ALLOWED_ORIGINS, got %v", err)
	}
}

func TestLoadInvalidPort(t *testing.T) {
	t.Setenv("ALLOWED_ORIGINS", "http://localhost:5173")
	for _, port := range []string{"abc", "70000"} {
		t.Setenv("PORT", port)
		_, err := Load()
		if err == nil || !strings.Contains(err.Error(), "PORT") {
			t.Fatalf("PORT=%s: expected error naming PORT, got %v", port, err)
		}
	}
}

func TestLoadDefaults(t *testing.T) {
	t.Setenv("ALLOWED_ORIGINS", "http://a.test, http://b.test")
	t.Setenv("PORT", "")
	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Port != "8080" {
		t.Fatalf("port = %q, want 8080", cfg.Port)
	}
	if len(cfg.AllowedOrigins) != 2 || cfg.AllowedOrigins[1] != "http://b.test" {
		t.Fatalf("origins = %v", cfg.AllowedOrigins)
	}
}
