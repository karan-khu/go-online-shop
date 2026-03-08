package inventory

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type InventoryHttpHandler interface {
	Listing(pctx *echo.Context) error
}

type InventoryUsecase interface {
	Listing(userId string) ([]*InventoryListing, error)
}

type InventoryRepository interface {
	Listing(userId string) ([]*QueryInventoryItemEntity, error)
	Filling(tx *gorm.DB, userId string, itemId int, qty int) ([]*InventoryEntity, error)
	Removing(tx *gorm.DB, userId string, itemId int, limit int) error
	UserItemCount(userId string, itemId int) int
}
