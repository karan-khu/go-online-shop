package balance

type CoinTopUpRequest struct {
	UserId string
	Amount int `json:"amount" validate:"required,gt=0"`
}

type UserCoinDisplay struct {
	UserId string `json:"user_id"`
	Coin   int    `json:"coin"`
}
