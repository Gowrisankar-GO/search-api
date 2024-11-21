package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name == "" {

		http.Error(w, "name query parameter is required", http.StatusBadRequest)

		return
	}

	w.Write([]byte(name))
}


func TestSearchAPI(t *testing.T) {

	req, err := http.NewRequest("GET", "/search?name=john", nil)

	if err != nil {
		t.Fatal("failed to create request:", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(searchHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {

		t.Errorf("expected status code 200, got %v", rr.Code)
	}

	expected := "john"

	if rr.Body.String() != expected {
		
		t.Errorf("expected body %v, got %v", expected, rr.Body.String())
	}
}
