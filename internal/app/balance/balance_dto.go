package balance

type CoinTopUpRequest struct {
	UserId string
	Amount int `json:"amount" validate:"required,gt=0"`
}
