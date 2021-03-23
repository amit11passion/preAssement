package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// To test TestAdd Functions
func TestAddTenEndpoint(t *testing.T) {
	request, err := http.NewRequest("GET", "/17", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"ip1": "17",
	}
	request = mux.SetURLVars(request, vars)
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(AddTen)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var data map[string]interface{}
	json.Unmarshal([]byte(response.Body.Bytes()), &data)
	var expected float64 = 27
	total := data["Sum"]
	if total != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			total, expected)
	}
}

// To test SumTwoVal Functions
func TestSumTwoValEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/17/10", nil)
	vars := map[string]string{
		"ip1": "17",
		"ip2": "10",
	}
	request = mux.SetURLVars(request, vars)
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(SumTwoVal)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	var data map[string]interface{}
	json.Unmarshal([]byte(response.Body.Bytes()), &data)
	// fmt.Println("data value", data)
	var expected float64 = 27
	total := data["Sum"]
	if total != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}
