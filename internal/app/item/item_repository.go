package item

import (
	"errors"
	"log/slog"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/internal/infra/database"
	"github.com/karan-khu/go-online-shop/internal/infra/database/models"
)

type itemRepositoryImpl struct {
	logger *slog.Logger
	db     database.Database
}

func NewItemRepository(logger *slog.Logger, conf *config.Config) ItemRepository {
	return &itemRepositoryImpl{
		logger: logger,
		db:     conf.GetDb("main"),
	}
}

func (r *itemRepositoryImpl) Listing(limit int, page int, searchText string) ([]*ItemEntity, int64, error) {
	itemRecords := make([]*models.ItemRecord, 0)

	query := r.db.Connect().Model(&models.ItemRecord{}).Where("ActiveStatus = ?", "AVAILABLE")
	if searchText != "" {
		searchText := "%" + searchText + "%"
		query = query.Where("Name LIKE ? OR Description LIKE ?", searchText, searchText)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		r.logger.Error("failed to get item total", "error", err)
		return nil, 0, err
	}

	query = query.Offset((page - 1) * limit).Limit(limit)
	if err := query.Find(&itemRecords).Error; err != nil {
		r.logger.Error("failed to get item listing", "error", err)
		return nil, 0, err
	}

	results := make([]*ItemEntity, 0)
	copier.Copy(&results, &itemRecords)
	return results, total, nil
}

func (r *itemRepositoryImpl) FindById(itemId int) (*ItemEntity, error) {
	itemRecord := new(models.ItemRecord)
	if err := r.db.Connect().Where("ItemId = ?", itemId).First(itemRecord).Error; err != nil {
		r.logger.Error("failed to find item by id", "error", err)
		return nil, errors.New("item not found")
	}

	result := new(ItemEntity)
	copier.Copy(result, itemRecord)
	return result, nil
}

func (r *itemRepositoryImpl) FindExists(itemId int) bool {
	var count int64
	if err := r.db.Connect().Model(&models.ItemRecord{}).Where("ItemId = ?", itemId).Count(&count).Error; err != nil {
		r.logger.Error("failed to find item exists", "error", err)
		return false
	}
	return count > 0
}

func (r *itemRepositoryImpl) Create(item *ItemEntity) (*ItemEntity, error) {
	newItem := new(models.ItemRecord)
	copier.Copy(newItem, item)

	itemRecord := new(models.ItemRecord)
	if err := r.db.Connect().Create(newItem).Scan(itemRecord).Error; err != nil {
		r.logger.Error("failed to create item", "error", err)
		return nil, err
	}

	result := new(ItemEntity)
	copier.Copy(result, itemRecord)
	return result, nil
}

func (r *itemRepositoryImpl) Edit(item *ItemEntity) (*ItemEntity, error) {
	updateItem := new(models.ItemRecord)
	copier.Copy(updateItem, item)

	if err := r.db.Connect().Where("ItemId = ?", item.ItemId).Updates(updateItem).Error; err != nil {
		r.logger.Error("failed to edit item", "error", err)
		return nil, err
	}

	itemRecord := new(models.ItemRecord)
	if err := r.db.Connect().Where("ItemId = ?", item.ItemId).First(itemRecord).Error; err != nil {
		r.logger.Error("failed to find item after edit", "error", err)
		return nil, err
	}

	result := new(ItemEntity)
	copier.Copy(result, itemRecord)
	return result, nil
}

func (r *itemRepositoryImpl) Archive(itemId int) error {
	if err := r.db.Connect().Model(&models.ItemRecord{}).Where("ItemId = ?", itemId).Update("ActiveStatus", "UNAVAILABLE").Error; err != nil {
		r.logger.Error("failed to archive item", "error", err)
		return err
	}
	return nil
}

func (r *itemRepositoryImpl) Begin() *gorm.DB {
	return r.db.Begin()
}

func (r *itemRepositoryImpl) Commit(tx *gorm.DB) error {
	return r.db.Commit(tx)
}

func (r *itemRepositoryImpl) Rollback(tx *gorm.DB) error {
	return r.db.Rollback(tx)
}
