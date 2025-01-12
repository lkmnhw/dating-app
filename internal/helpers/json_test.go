package helpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteJSONResponse(t *testing.T) {
	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Sample data to write as JSON
	data := map[string]string{"message": "success"}

	// Call the WriteJSONResponse function
	WriteJSONResponse(rr, http.StatusOK, data)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the Content-Type header
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Check the response body
	var response map[string]string
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, data, response)
}

func TestWriteJSONResponse_Error(t *testing.T) {
	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the WriteJSONResponse function with a nil data
	WriteJSONResponse(rr, http.StatusInternalServerError, nil)

	// Check the status code
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Check the Content-Type header
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Check the response body
	var response interface{}
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Nil(t, response)
}
