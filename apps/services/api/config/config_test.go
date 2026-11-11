package config

import (
	"reflect"
	"strings"
	"testing"
)

func TestLoadRequiresAllowedOrigins(t *testing.T) {
	for _, v := range []string{"", "   "} {
		t.Setenv("ALLOWED_ORIGINS", v)
		_, err := Load()
		if err == nil || !strings.Contains(err.Error(), "ALLOWED_ORIGINS") {
			t.Fatalf("value %q: want error naming ALLOWED_ORIGINS, got %v", v, err)
		}
	}
}

func TestLoadRejectsBadPort(t *testing.T) {
	t.Setenv("ALLOWED_ORIGINS", "http://localhost:5173")
	for _, p := range []string{"abc", "0", "70000"} {
		t.Setenv("PORT", p)
		_, err := Load()
		if err == nil || !strings.Contains(err.Error(), "PORT") {
			t.Fatalf("PORT %q: want error naming PORT, got %v", p, err)
		}
	}
}

func TestLoadDefaultsAndSplits(t *testing.T) {
	t.Setenv("ALLOWED_ORIGINS", " http://a.test , http://b.test ")
	t.Setenv("PORT", "")
	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Port != "8080" {
		t.Fatalf("port = %q", cfg.Port)
	}
	if want := []string{"http://a.test", "http://b.test"}; !reflect.DeepEqual(cfg.AllowedOrigins, want) {
		t.Fatalf("origins = %v", cfg.AllowedOrigins)
	}
}
