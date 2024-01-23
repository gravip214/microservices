// handlers_test.go
package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetTopTrackInfo(t *testing.T) {
	// Create a new request with a dummy region
	req, err := http.NewRequest("GET", "/top-track/usa", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Create a mux router and attach the handler
	router := mux.NewRouter()
	router.HandleFunc("/top-track/{region}", GetTopTrackInfo)

	// Serve the request to the recorder
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

}
