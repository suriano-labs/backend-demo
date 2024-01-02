package entity

type Credit struct {
	UserID  string `json:"user_id"`
	ApiPath string `json:"api_path"`
	Credit  int    `json:"credit"`
}
