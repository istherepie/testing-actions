package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	mux := New("TESTINSTANCE-1")

	mux.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Incorrect HTTP status ,got %v want %v", status, http.StatusOK)
	}

	expected := `{"instance":"TESTINSTANCE-1FAIL"}`

	output := recorder.Body.String()

	if output != expected {
		t.Errorf("Response body error: got %v want %v", output, expected)
	}
}
