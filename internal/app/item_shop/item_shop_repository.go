package itemshop

import (
	"log/slog"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/database"
	"github.com/karan-khu/go-online-shop/pkg/database/models"
)

type itemShopRepositoryImpl struct {
	db     database.Database
	logger *slog.Logger
}

func NewItemShopRepository(logger *slog.Logger, conf *config.Config) ItemShopRepository {
	return &itemShopRepositoryImpl{
		db: conf.GetDb("main"),
	}
}

func (r *itemShopRepositoryImpl) Listing(limit int, page int, searchText string) (results []*ItemShopEntity, total int64, err error) {
	itemRecords := make([]*models.ItemRecord, 0)

	query := r.db.Connect().Model(&models.ItemRecord{}).Where("ActiveStatus = ?", "AVAILABLE")
	if searchText != "" {
		searchText := "%" + searchText + "%"
		query = query.Where("Name LIKE ? OR Description LIKE ?", searchText, searchText)
	}

	if err = query.Count(&total).Error; err != nil {
		r.logger.Error("failed to get item total", "error", err)
		return nil, 0, err
	}

	query = query.Offset((page - 1) * limit).Limit(limit)
	if err = query.Find(&itemRecords).Error; err != nil {
		r.logger.Error("failed to get item listing", "error", err)
		return nil, 0, err
	}

	results = make([]*ItemShopEntity, 0)
	copier.Copy(&results, &itemRecords)
	return results, total, nil
}

func (r *itemShopRepositoryImpl) Begin() *gorm.DB {
	return r.db.Connect().Begin()
}

func (r *itemShopRepositoryImpl) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *itemShopRepositoryImpl) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}
