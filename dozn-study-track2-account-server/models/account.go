package models

type Account struct {
	UserId      int     `json:"userId"`
	AccountType string  `json:"accountType"`
	Interest    float64 `json:"interest"`
}
