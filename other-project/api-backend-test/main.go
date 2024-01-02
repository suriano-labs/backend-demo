package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type HealthCheck struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func init() {
	fmt.Println("API-Client-Test: ON")
}

func main() {
	http.HandleFunc("/health-check", HealthCheckHandler)
	http.HandleFunc("/user", UserHandler)
	http.HandleFunc("/user-vagalume", UserVagalumeHandler)
	http.ListenAndServe(":8000", nil)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	healthCheck := HealthCheck{Status: "ok", Message: "Hello"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthCheck)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		user := User{Name: "João", Age: "25"}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func UserVagalumeHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "João", Age: "25"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
