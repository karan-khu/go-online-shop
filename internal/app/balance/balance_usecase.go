package balance

type BalanceUsecaseImpl struct {
	balanceRepo BalanceRepository
}

func NewBalanceUsecase(balanceRepo BalanceRepository) BalanceUsecase {
	return &BalanceUsecaseImpl{
		balanceRepo: balanceRepo,
	}
}

func (u *BalanceUsecaseImpl) CoinAdding(req *CoinTopUpRequest) (*UserBalanceEntity, error) {
	balance := &UserBalanceEntity{
		UserId: req.UserId,
		Amount: req.Amount,
		Status: "TOPUP",
	}

	newBalance, err := u.balanceRepo.CoinAdd(balance)
	if err != nil {
		return nil, err
	}

	return newBalance, nil
}

func (u *BalanceUsecaseImpl) CoinShowBalance(userId string) *UserCoinDisplay {
	coin, err := u.balanceRepo.CoinShow(userId)
	if err != nil {
		return &UserCoinDisplay{
			UserId: userId,
			Coin:   0,
		}
	}

	return coin
}
