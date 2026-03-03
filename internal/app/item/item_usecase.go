package item

import (
	"log/slog"
	"math"

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

func (u *itemUsecaseImpl) ItemList(req *RequestItemFilter) (*ResponseItemList, error) {
	list, total, err := u.itemRepo.Listing(req)
	if err != nil {
		u.logger.Error("failed to get item list", "error", err)
		return nil, err
	}

	itemModels := make([]*ItemModel, len(list))
	for i, item := range list {
		itemModels[i] = item.ToModel(u.host)
	}

	return &ResponseItemList{
		Items: itemModels,
		Pagination: &Pagination{
			Total:      total,
			TotalPages: int(math.Ceil(float64(total) / float64(req.Limit))),
		},
	}, nil
}

func (u *itemUsecaseImpl) CreateItem(req *RequestItemCreate) (*ItemModel, error) {
	item, err := u.itemRepo.Create(req.ToEntity())
	if err != nil {
		u.logger.Error("failed to create item", "error", err)
		return nil, err
	}

	return item.ToModel(u.host), nil
}
