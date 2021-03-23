package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// result: Struct to store final sum value
type result struct {
	Sum int64 `json:"sum`
}

// setting up routes
func initializeRouter() {
	godotenv.Load()
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/{ip1:[0-9]+}", AddTen).Methods("GET")
	router.HandleFunc("/{ip1:[0-9]+}/{ip2:[0-9]+}", SumTwoVal).Methods("GET")
	PORT := fmt.Sprintf("%s", port)
	http.ListenAndServe(":"+PORT, router)

}

//AddTen : fun that add 10 to input and retun result
func AddTen(w http.ResponseWriter, r *http.Request) {
	var resultSum result
	ip1 := mux.Vars(r)["ip1"]
	w.Header().Set("Content-type", "application/json")
	intIp1, _ := strconv.ParseInt(ip1, 10, 64)
	sm := intIp1 + 10
	resultSum.Sum = sm
	json.NewEncoder(w).Encode(resultSum)

}

// SumTwoVal: func to add two input and return result
func SumTwoVal(w http.ResponseWriter, r *http.Request) {
	var resultSum result
	ip1 := mux.Vars(r)["ip1"]
	ip2 := mux.Vars(r)["ip2"]
	w.Header().Set("Content-type", "application/json")
	intIp1, _ := strconv.ParseInt(ip1, 10, 64)
	intIp2, _ := strconv.ParseInt(ip2, 10, 64)
	resultSum.Sum = intIp1 + intIp2
	json.NewEncoder(w).Encode(resultSum)

}

func main() {
	fmt.Print("Starting Api on port: 5000")
	initializeRouter()
}
