package handlers

import (
	"bill-manager/internal/entity"
	"bill-manager/internal/infra/database"
	"encoding/json"
	"net/http"
)

type CreditHandler struct {
	CreditDB database.CreditInterface
}

type Error struct {
	Message string `json:"message"`
}

func NewCreditHandler(creditDB database.CreditInterface) *CreditHandler {
	return &CreditHandler{
		CreditDB: creditDB,
	}
}

func (c *CreditHandler) CreditCheck(w http.ResponseWriter, r *http.Request) {
	bill := entity.Bill{}

	err := json.NewDecoder(r.Body).Decode(&bill)
	if err != nil {
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	credit, err := c.CreditDB.FindByUserApi(bill.UserID, bill.ApiPath)
	// if err != nil {
	// 	error := Error{Message: err.Error()}
	// 	json.NewEncoder(w).Encode(error)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }

	if credit.Credit > 0 {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized) //401
		return
	}
}
