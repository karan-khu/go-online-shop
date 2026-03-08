package balance

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type BalanceHttpHandler interface {
	CoinTopUp(pctx *echo.Context) error
	CoinShowBalance(pctx *echo.Context) error
}

type BalanceUsecase interface {
	CoinAdding(req *CoinTopUpRequest) (*UserBalanceEntity, error)
	CoinShowBalance(userId string) *UserCoinDisplay
}

type BalanceRepository interface {
	CoinAdd(tx *gorm.DB, req *UserBalanceEntity) (*UserBalanceEntity, error)
	CoinShow(userId string) (*UserCoinDisplay, error)
}
