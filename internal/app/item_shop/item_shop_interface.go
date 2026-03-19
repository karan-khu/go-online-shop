package itemshop

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type ItemShopHttpHandler interface {
	GetAll(c *echo.Context) error
	Buying(c *echo.Context) error
	Selling(c *echo.Context) error
}

type ItemShopUsecase interface {
	ItemList(filter *RequestItemFilter) (*ResponseItemList, error)
	Buying(req *RequestItemBuying) error
	Selling(req *RequestItemSelling) error
}

type ItemShopRepository interface {
	Listing(limit int, page int, searchText string) ([]*ItemShopEntity, int64, error)
	Begin() *gorm.DB
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
}
