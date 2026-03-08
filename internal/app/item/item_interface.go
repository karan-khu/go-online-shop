package item

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type ItemHttpHandler interface {
	GetAll(c *echo.Context) error
	Create(c *echo.Context) error
	Edit(c *echo.Context) error
	Delete(c *echo.Context) error
	Buying(c *echo.Context) error
	Selling(c *echo.Context) error
}

type ItemUsecase interface {
	ItemList(filter *RequestItemFilter) (*ResponseItemList, error)
	CreateItem(req *RequestItemCreate) (*ItemModel, error)
	EditItem(req *RequestItemEdit) (*ItemModel, error)
	DeleteItem(itemId int) error
	Buying(req *RequestItemBuying) error
	Selling(req *RequestItemSelling) error
}

type ItemRepository interface {
	Listing(filter *RequestItemFilter) ([]*ItemEntity, int, error)
	FindById(itemId int) (*ItemEntity, error)
	FindExists(itemId int) bool
	Create(item *ItemEntity) (*ItemEntity, error)
	Edit(item *ItemEntity) (*ItemEntity, error)
	Archive(itemId int) error
	Begin() *gorm.DB
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
}
