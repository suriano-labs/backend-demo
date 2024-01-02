package main

import (
	"bill-manager/internal/infra/database"
	"bill-manager/internal/infra/webserver/handlers"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	fmt.Println("API-Billing-Test: ON")
}

func main() {

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}

	creditDB := database.NewCreditDB(db)
	creditHandler := handlers.NewCreditHandler(creditDB)

	billDB := database.NewBillDB(db)
	billHandler := handlers.NewBillHandler(*billDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health-check", HealthCheckHandler)
	r.Post("/credit-check", creditHandler.CreditCheck)
	r.Post("/bill-charge", billHandler.BillCharge)

	http.ListenAndServe(":8001", r)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	healthCheck := "Living and kick"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthCheck)
}
