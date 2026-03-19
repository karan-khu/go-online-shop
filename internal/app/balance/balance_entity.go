package balance

import "time"

type UserBalanceEntity struct {
	BalanceId int
	UserId    string
	Amount    int
	Status    string
	CreatedAt time.Time
}

type UserCoin struct {
	UserId   string `json:"user_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Coin     int    `json:"coin"`
}
