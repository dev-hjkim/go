type Account struct {
	AccountId string `json:"accountId" validate:"required"`
	UserId string `json:"userId" validate:"required"`
	Username string `json:"name" validate:"required"`
}