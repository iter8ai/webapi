package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	cases := []struct {
		method, path string
		ok           bool
	}{
		{http.MethodGet, "/healthz", true},
		{http.MethodPost, "/healthz", false},
		{http.MethodGet, "/nope", false},
	}
	for _, c := range cases {
		rec := httptest.NewRecorder()
		newRouter().ServeHTTP(rec, httptest.NewRequest(c.method, c.path, nil))
		if (rec.Code == http.StatusOK) != c.ok {
			t.Errorf("%s %s = %d, ok want %v", c.method, c.path, rec.Code, c.ok)
		}
	}
}
