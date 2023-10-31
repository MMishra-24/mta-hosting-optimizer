package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHostnamesWithActiveIPs(t *testing.T) {

	r, _ := InitializeRouter()
	os.Setenv("X", "")
	// Create a test request to the /hostnames endpoint.
	req, err := http.NewRequest("GET", "/hostnames", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	assert.JSONEq(t, `{"hostnames":["mta-prod-1", "mta-prod-3"]}`, rr.Body.String())
}

func TestGetHostnamesWithActiveIPsValidations(t *testing.T) {
	r, _ := InitializeRouter()

	//Change the envrionment variable
	_ = os.Setenv("X", "1E")

	req, err := http.NewRequest("GET", "/hostnames", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
