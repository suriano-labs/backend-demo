package handlers

import (
	"bill-manager/internal/entity"
	"bill-manager/internal/infra/database"
	"encoding/json"
	"net/http"
	"time"
)

type BillHandler struct {
	BillDB database.Bill
}

func NewBillHandler(db database.Bill) *BillHandler {
	return &BillHandler{
		BillDB: db,
	}
}

func (b *BillHandler) BillCharge(w http.ResponseWriter, r *http.Request) {
	payload := entity.Bill{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if payload.ApiPath == "" || payload.UserID == "" {
		// error := Error{Message: ""}
		// json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bill, err := b.BillDB.Create(payload.UserID, payload.ApiPath, "Ok", time.Now())
	if err != nil {
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(bill)
	if err != nil {
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
