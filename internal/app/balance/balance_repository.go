package balance

import (
	"errors"
	"log/slog"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/database"
	"github.com/karan-khu/go-online-shop/pkg/database/models"
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

func (r *BalanceRepositoryImpl) CoinAdd(tx *gorm.DB, req *UserBalanceEntity) (result *UserBalanceEntity, err error) {
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

	result = new(UserBalanceEntity)
	copier.Copy(result, balanceRecord)
	return result, nil
}

func (r *BalanceRepositoryImpl) CoinShow(userId string) (result *UserCoin, err error) {
	result = new(UserCoin)
	if err := r.db.Connect().Model(&models.UserBalanceRecord{}).
		Where("GOST_UserBalances.UserId = ?", userId).
		Select("UserId, user.FullName, user.Email, SUM(Amount) AS Coin").
		Joins("JOIN GOST_Users AS user ON user.UserId = GOST_UserBalances.UserId").
		Group("GOST_UserBalances.UserId").
		Scan(result).Error; err != nil {
		r.logger.Error("failed to show coin", "error", err)
		return nil, errors.New("failed to get coin")
	}

	return result, nil
}
