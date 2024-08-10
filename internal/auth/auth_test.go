package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header(make(map[string][]string))
	headers.Set("Authorization", "ApiKey caligastia")
	got, err := GetAPIKey(headers)
	want := "caligastia"
	if err != nil {
		t.Fatalf("GetAPIKey returned unexpected error!")
	}

	if got != want {
		t.Fatalf("did not get api key back!")
	}

	headers = http.Header(make(map[string][]string))
	_, err = GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected ErrNoAuthHeaderIncluded")
	}

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("an unexpected error was returned from GetAPIKey")
	}
}
