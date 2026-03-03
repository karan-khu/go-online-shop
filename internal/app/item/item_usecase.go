package item

import (
	"log/slog"

	"github.com/karan-khu/go-online-shop/config"
)

type itemUsecaseImpl struct {
	logger   *slog.Logger
	itemRepo ItemRepository
	host     string
}

func NewItemUsecase(logger *slog.Logger, itemRepo ItemRepository, conf *config.Config) ItemUsecase {
	return &itemUsecaseImpl{
		logger:   logger,
		itemRepo: itemRepo,
		host:     conf.Env.APP_HOST,
	}
}

func (u *itemUsecaseImpl) ItemList() ([]*ItemModel, error) {
	list, err := u.itemRepo.Listing()
	if err != nil {
		u.logger.Error("failed to get item list", "error", err)
		return nil, err
	}

	itemModels := make([]*ItemModel, len(list))
	for i, item := range list {
		itemModels[i] = item.ToModel(u.host)
	}

	return itemModels, nil
}
