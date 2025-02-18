package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "")

	_, err := GetAPIKey(header)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf(`Expected ErrNoAuthHeaderIncluded, got %v`, err)
	}

	header = http.Header{}
	header.Add("Authorization", "Bearer 1234")

	_, err = GetAPIKey(header)
	if err == nil {
		t.Fatalf(`Expected error, got %v`, err)
	}

	header = http.Header{}
	header.Add("Authorization", "ApiKey 1234")
	apiKey, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf(`Expected no error, got %v`, err)
	}
	if apiKey != "1234" {
		t.Fatalf(`Expected "1234", got %s"`, apiKey)
	}
}
