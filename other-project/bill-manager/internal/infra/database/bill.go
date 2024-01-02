package database

import (
	"bill-manager/internal/entity"
	"database/sql"
	"time"
)

type Bill struct {
	db            *sql.DB
	UserID        string    `json:"user_id"`
	ApiPath       string    `json:"api_path"`
	Status        string    `json:"status"`
	ProcessedDate time.Time `json:"processed_date"`
}

func NewBillDB(db *sql.DB) *Bill {
	return &Bill{
		db: db,
	}
}

func (b *Bill) Create(userID string, apiPath string, status string, date time.Time) (*entity.Bill, error) {

	query := "INSERT INTO bills (user_id, api_path, status, processed_date) VALUES ($1, $2, $3, $4)"

	_, err := b.db.Exec(query, userID, apiPath, status, date)
	if err != nil {
		return nil, err
	}

	return &entity.Bill{
		UserID:        userID,
		ApiPath:       apiPath,
		Status:        status,
		ProcessedDate: date,
	}, nil
}
