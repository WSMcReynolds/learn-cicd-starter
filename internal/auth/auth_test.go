package auth

import (
	"net/http"
	"testing"
)

func TestValidGetAPIKey(t *testing.T) {
	// Test valid request
	header := http.Header{}
	header.Add("Authorization", "ApiKey 12345")

	_, err := GetAPIKey(header)

	if err != nil {
		t.Errorf(`APIKey was not returned... Auth Header: %s`, header.Get("Authorization"))
	}

}

func TestInvalidGetAPIKey(t *testing.T) {
	// Test invalid header (no auth)
	badHeader := http.Header{}
	badHeader.Add("Content-Type", "application/html")
	_, err := GetAPIKey(badHeader)

	if err == nil {
		t.Errorf(`Invalid header with no authorization didn't fail. Header: %s`, badHeader.Get("Authorization"))
	}
}
