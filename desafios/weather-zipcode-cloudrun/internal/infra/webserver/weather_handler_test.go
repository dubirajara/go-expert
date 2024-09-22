package webserver

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:8000/weather/"

func TestZipCodeValid(t *testing.T) {
	resp, err := http.Get(baseURL + "?zipcode=78075588")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestZipCodeInValid(t *testing.T) {
	resp, err := http.Get(baseURL + "?zipcode=707588")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}

func TestZipCodeNotFind(t *testing.T) {
	resp, err := http.Get(baseURL + "?zipcode=88075-588")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
