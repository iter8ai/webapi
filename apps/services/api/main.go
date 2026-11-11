package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/iter8ai/webapi/apps/services/api/config"
)

func newRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	return r
}

func main() {
	// godotenv does not override variables already set, so the real environment wins.
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := http.ListenAndServe(":"+cfg.Port, newRouter()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
