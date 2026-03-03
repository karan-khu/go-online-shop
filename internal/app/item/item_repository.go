package item

import (
	"log/slog"

	"gorm.io/gorm"

	"github.com/karan-khu/go-online-shop/config"
)

type itemRepositoryImpl struct {
	logger *slog.Logger
	db     *gorm.DB
}

func NewItemRepository(logger *slog.Logger, conf *config.Config) ItemRepository {
	return &itemRepositoryImpl{
		logger: logger,
		db:     conf.GetDb("main"),
	}
}

func (r *itemRepositoryImpl) Listing(req *RequestItemFilter) (list []*ItemEntity, err error) {
	query := r.db.Model(&ItemEntity{})
	if req.SearchText != "" {
		searchText := "%" + req.SearchText + "%"
		query = query.Where("Name LIKE ? OR Description LIKE ?", searchText, searchText)
	}

	err = query.Find(&list).Error
	if err != nil {
		r.logger.Error("failed to get item listing", "error", err)
		return nil, err
	}

	return list, nil
}
