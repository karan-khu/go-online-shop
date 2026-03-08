package purchase

import (
	"errors"
	"log/slog"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/database"
)

type purchaseRepositoryImpl struct {
	logger *slog.Logger
	db     database.Database
}

func NewPurchaseRepository(logger *slog.Logger, conf *config.Config) PurchaseRepository {
	return &purchaseRepositoryImpl{
		logger: logger,
		db:     conf.GetDb("main"),
	}
}

func (r *purchaseRepositoryImpl) Listing(userId string) ([]*PurchaseHistoryEntity, error) {
	listPurchase := make([]*PurchaseHistoryEntity, 0)
	if err := r.db.Connect().Where("BuyerId = ? AND ActiveStatus = ?", userId, "AVAILABLE").Find(&listPurchase).Error; err != nil {
		r.logger.Error("failed to list purchase", "error", err)
		return nil, errors.New("failed to list purchase")
	}

	return listPurchase, nil
}

func (r *purchaseRepositoryImpl) Create(purchase *PurchaseHistoryEntity) (*PurchaseHistoryEntity, error) {
	newPurchase := new(PurchaseHistoryEntity)
	if err := r.db.Connect().Create(purchase).Scan(newPurchase).Error; err != nil {
		r.logger.Error("failed to create purchase", "error", err)
		return nil, errors.New("failed to create purchase")
	}
	return newPurchase, nil
}
