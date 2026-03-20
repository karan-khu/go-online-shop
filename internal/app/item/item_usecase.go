package item

import (
	"errors"
	"log/slog"

	"github.com/karan-khu/go-online-shop/internal/upload"
)

type itemUsecaseImpl struct {
	logger       *slog.Logger
	itemRepo     ItemRepository
	imageBuilder upload.ImageBuilder
}

func NewItemUsecase(logger *slog.Logger, itemRepo ItemRepository, imageBuilder upload.ImageBuilder) ItemUsecase {
	return &itemUsecaseImpl{
		logger:       logger,
		itemRepo:     itemRepo,
		imageBuilder: imageBuilder,
	}
}

func (u *itemUsecaseImpl) CreateItem(req *RequestItemCreate, adminId string) (*ItemEntity, error) {
	entity := req.ToEntity()
	entity.AdminId = adminId
	item, err := u.itemRepo.Create(entity)
	if err != nil {
		u.logger.Error("failed to create item", "error", err)
		return nil, err
	}

	return &ItemEntity{
		ItemId:      item.ItemId,
		Name:        item.Name,
		Description: item.Description,
		Picture:     u.imageBuilder.Build(item.Picture),
		Price:       item.Price,
	}, nil
}

func (u *itemUsecaseImpl) EditItem(req *RequestItemEdit) (*ItemEntity, error) {
	item, err := u.itemRepo.Edit(req.ToEntity())
	if err != nil {
		u.logger.Error("failed to edit item", "error", err)
		return nil, err
	}

	return &ItemEntity{
		ItemId:      item.ItemId,
		Name:        item.Name,
		Description: item.Description,
		Picture:     u.imageBuilder.Build(item.Picture),
		Price:       item.Price,
	}, nil
}

func (u *itemUsecaseImpl) DeleteItem(itemId int) error {
	if !u.itemRepo.FindExists(itemId) {
		return errors.New("item does not exist")
	}

	err := u.itemRepo.Archive(itemId)
	if err != nil {
		u.logger.Error("failed to delete item", "error", err)
		return err
	}

	return nil
}
