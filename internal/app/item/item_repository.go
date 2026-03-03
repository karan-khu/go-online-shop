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

func (r *itemRepositoryImpl) Listing() (list []*ItemEntity, err error) {
	err = r.db.Model(&ItemEntity{}).Find(&list).Error
	if err != nil {
		r.logger.Error("failed to get item listing", "error", err)
		return nil, err
	}

	return list, nil
}
