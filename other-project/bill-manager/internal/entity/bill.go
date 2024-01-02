package entity

import "time"

type Bill struct {
	UserID        string    `json:"user_id"`
	ApiPath       string    `json:"api_path"`
	Status        string    `json:"status"`
	ProcessedDate time.Time `json:"processed_date"`
}

func NewBill(userID, apiPath, status string, processedDate time.Time) *Bill {
	return &Bill{
		UserID:        userID,
		ApiPath:       apiPath,
		Status:        status,
		ProcessedDate: processedDate,
	}
}
