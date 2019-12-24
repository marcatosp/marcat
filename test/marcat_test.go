package marcat

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/minIddamal/marcat"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(marcat.HealthHandler())



}