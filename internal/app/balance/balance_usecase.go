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

	newBalance, err := u.balanceRepo.CoinAdd(nil, balance)
	if err != nil {
		return nil, err
	}

	return newBalance, nil
}

func (u *BalanceUsecaseImpl) CoinShowBalance(userId string) *UserCoin {
	coin, err := u.balanceRepo.CoinShow(userId)
	if err != nil {
		return &UserCoin{
			UserId: userId,
			Coin:   0,
		}
	}

	return coin
}
