package inventory

import (
	"errors"
	"log/slog"

	"gorm.io/gorm"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/internal/infra/database"
)

type inventoryRepositoryImpl struct {
	logger *slog.Logger
	db     database.Database
}

func NewInventoryRepository(logger *slog.Logger, conf *config.Config) InventoryRepository {
	return &inventoryRepositoryImpl{
		logger: logger,
		db:     conf.GetDb("main"),
	}
}

func (r *inventoryRepositoryImpl) Listing(userId string) ([]*InventoryEntity, error) {
	listInventory := make([]*InventoryEntity, 0)
	if err := r.db.Connect().
		Raw(`
			SELECT 
				i.*,
				item.Name,
				item.Description,
				item.Picture
			FROM GOST_Items AS item
			INNER JOIN GOST_Inventories AS i ON i.ItemId = item.ItemId
			WHERE i.UserId = ?
				AND i.ActiveStatus = 'AVAILABLE'
		`, userId).
		Scan(&listInventory).
		Error; err != nil {
		r.logger.Error("failed to listing inventory", "error", err)
		return nil, errors.New("failed to listing inventory")
	}

	return listInventory, nil
}

func (r *inventoryRepositoryImpl) Filling(tx *gorm.DB, userId string, itemId int, qty int) ([]*InventoryEntity, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	inventoryList := make([]*InventoryEntity, 0)
	for range qty {
		inventoryList = append(inventoryList, &InventoryEntity{
			UserId: userId,
			ItemId: itemId,
		})
	}

	if err := conn.CreateInBatches(inventoryList, len(inventoryList)).Error; err != nil {
		r.logger.Error("InventoryRepository.Filling", "error", err)
		return nil, errors.New("failed to fill inventory")
	}

	return inventoryList, nil
}

func (r *inventoryRepositoryImpl) Removing(tx *gorm.DB, userId string, itemId int, limit int) error {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	if err := conn.Exec(`
		UPDATE TOP(?) i
		SET ActiveStatus = 'UNAVAILABLE'
		FROM GOST_Inventories i
		WHERE UserId = ? AND ItemId = ? AND ActiveStatus = 'AVAILABLE'
	`, limit, userId, itemId).Error; err != nil {
		r.logger.Error("InventoryRepository.Removing", "error", err)
		return errors.New("failed to remove inventory")
	}
	return nil
}

func (r *inventoryRepositoryImpl) UserItemCount(userId string, itemId int) int {
	var count int64
	if err := r.db.Connect().
		Model(&InventoryEntity{}).
		Where("UserId = ? AND ItemId = ? AND ActiveStatus = ?", userId, itemId, "AVAILABLE").
		Count(&count).
		Error; err != nil {
		r.logger.Error("InventoryRepository.UserItemCount", "error", err)
		return -1
	}

	return int(count)
}
