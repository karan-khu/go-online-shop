package purchase

import (
	"errors"
	"log/slog"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/database"
	"github.com/karan-khu/go-online-shop/pkg/database/models"
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

func (r *purchaseRepositoryImpl) Listing(userId string) (results []*PurchaseHistoryEntity, err error) {
	purchaseRecords := make([]*models.PurchaseHistoryRecord, 0)
	if err := r.db.Connect().Where("BuyerId = ? AND ActiveStatus = ?", userId, "AVAILABLE").Find(&purchaseRecords).Error; err != nil {
		r.logger.Error("failed to list purchase", "error", err)
		return nil, errors.New("failed to list purchase")
	}

	copier.Copy(results, purchaseRecords)
	return results, nil
}

func (r *purchaseRepositoryImpl) Create(tx *gorm.DB, purchase *PurchaseHistoryEntity) (result *PurchaseHistoryEntity, err error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	newPurchase := new(models.PurchaseHistoryRecord)
	copier.Copy(newPurchase, purchase)

	purchaseRecord := new(models.PurchaseHistoryRecord)
	if err := conn.Create(newPurchase).Scan(purchaseRecord).Error; err != nil {
		r.logger.Error("failed to create purchase", "error", err)
		return nil, errors.New("failed to create purchase")
	}

	copier.Copy(result, purchaseRecord)
	return result, nil
}
