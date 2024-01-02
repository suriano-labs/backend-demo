package database

import (
	"bill-manager/internal/entity"
	"time"
)

type BillInterface interface {
	Create(userID string, apiPath string, status string, date time.Time) (*entity.Bill, error)
}

type CreditInterface interface {
	FindByUserApi(userID, apiID string) (entity.Credit, error)
	InsertCredit(userID, apiID string, credits int) (entity.Credit, error)
}
