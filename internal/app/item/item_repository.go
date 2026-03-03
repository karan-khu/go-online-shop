package item

import (
	"log/slog"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/database"
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

func (r *itemRepositoryImpl) Listing(req *RequestItemFilter) ([]*ItemEntity, int, error) {
	list := make([]*ItemEntity, 0)

	query := r.db.Connect().Model(&ItemEntity{}).Where("ActiveStatus = ?", "AVAILABLE")
	if req.SearchText != "" {
		searchText := "%" + req.SearchText + "%"
		query = query.Where("Name LIKE ? OR Description LIKE ?", searchText, searchText)
	}

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		r.logger.Error("failed to get item total", "error", err)
		return nil, 0, err
	}

	query = query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	if err := query.Find(&list).Error; err != nil {
		r.logger.Error("failed to get item listing", "error", err)
		return nil, 0, err
	}
	return list, int(totalCount), nil
}
