package database

import (
	"bill-manager/internal/entity"
	"database/sql"
)

type Credit struct {
	db     *sql.DB
	UserID string `json:"user_id"`
	ApiID  string `json:"api_id"`
	Credit int    `json:"credit"`
}

func NewCreditDB(db *sql.DB) *Credit {
	return &Credit{
		db: db,
	}
}

func (c *Credit) FindByUserApi(userID, apiPath string) (entity.Credit, error) {
	var user_id, api_path string
	var credit int

	query := "SELECT user_id, api_path, credit FROM credits WHERE user_id=$1 and api_path=$2"

	row := c.db.QueryRow(query, userID, apiPath)
	err := row.Scan(&user_id, &api_path, &credit)
	if err != nil {
		return entity.Credit{}, err
	}

	return entity.Credit{
		UserID:  user_id,
		ApiPath: api_path,
		Credit:  credit,
	}, nil

}

func (c *Credit) InsertCredit(userID, apiPath string, credits int) (entity.Credit, error) {
	credit, err := c.FindByUserApi(userID, apiPath)
	if err != nil {
		return entity.Credit{}, err
	}

	var query string

	if credit.UserID == "" {
		query = "INSERT INTO credits (user_id, api_id, credit) VALUES ($1, $2, $3)"
	}

	query = "INSERT INTO credits (credit) VALUES ($3) WHERE user_id=$1 and api_id=$2"

	_, err = c.db.Exec(query, userID, apiPath, credits)
	if err != nil {
		return entity.Credit{}, err
	}

	return entity.Credit{
		UserID:  userID,
		ApiPath: apiPath,
		Credit:  credits,
	}, nil
}
