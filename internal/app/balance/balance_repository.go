package balance

import (
	"errors"
	"log/slog"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/internal/infra/database"
	"github.com/karan-khu/go-online-shop/internal/infra/database/models"
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

func (r *BalanceRepositoryImpl) CoinAdd(tx *gorm.DB, req *UserBalanceEntity) (*UserBalanceEntity, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}
	newBalance := new(models.UserBalanceRecord)
	copier.Copy(newBalance, req)

	balanceRecord := new(models.UserBalanceRecord)
	if err := conn.Create(newBalance).Scan(balanceRecord).Error; err != nil {
		r.logger.Error("failed to create balance", "error", err)
		return nil, errors.New("failed to create balance")
	}

	result := new(UserBalanceEntity)
	copier.Copy(result, balanceRecord)
	return result, nil
}

func (r *BalanceRepositoryImpl) CoinShow(userId string) (*UserCoin, error) {
	result := new(UserCoin)
	if err := r.db.Connect().Model(&models.UserBalanceRecord{}).
		Select("u.UserId, u.FullName, u.Email, SUM(GOST_UserBalances.Amount) AS Coin").
		Joins("RIGHT JOIN GOST_Users AS u ON u.UserId = GOST_UserBalances.UserId").
		Where("u.UserId = ?", userId).
		Group("u.UserId, u.FullName, u.Email").
		Scan(result).Error; err != nil {
		r.logger.Error("failed to show coin", "error", err)
		return nil, errors.New("failed to get coin")
	}

	return result, nil
}
