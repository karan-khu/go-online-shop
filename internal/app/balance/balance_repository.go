package balance

import (
	"log/slog"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/database"
)

type BalanceRepositoryImpl struct {
	logger *slog.Logger
	db     database.Database
}

func NewBalanceRepository(logger *slog.Logger, conf *config.Config) BalanceRepository {
	return &BalanceRepositoryImpl{
		logger: logger,
		db:     conf.GetDb("main"),
	}
}

func (r *BalanceRepositoryImpl) CoinAdd(req *UserBalanceEntity) (*UserBalanceEntity, error) {
	newBalance := new(UserBalanceEntity)
	if err := r.db.Connect().Create(req).Scan(newBalance).Error; err != nil {
		r.logger.Error("failed to create balance", "error", err)
		return nil, err
	}

	return newBalance, nil
}

func (r *BalanceRepositoryImpl) CoinShow(userId string) (*UserCoinDisplay, error) {
	coin := new(UserCoinDisplay)

	if err := r.db.Connect().Model(&UserBalanceEntity{}).
		Where("UserId = ?", userId).
		Select("UserId, SUM(Amount) AS Coin").
		Group("UserId").
		Scan(coin).Error; err != nil {
		r.logger.Error("failed to show coin", "error", err)
		return nil, err
	}

	return coin, nil
}
