package balance

import "github.com/labstack/echo/v5"

type BalanceHttpHandler interface {
	CoinTopUp(pctx *echo.Context) error
	CoinShowBalance(pctx *echo.Context) error
}

type BalanceUsecase interface {
	CoinAdding(req *CoinTopUpRequest) (*UserBalanceEntity, error)
	CoinShowBalance(userId string) *UserCoinDisplay
}

type BalanceRepository interface {
	CoinAdd(req *UserBalanceEntity) (*UserBalanceEntity, error)
	CoinShow(userId string) (*UserCoinDisplay, error)
}
