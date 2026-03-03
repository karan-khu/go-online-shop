package item

import "github.com/labstack/echo/v5"

type ItemHttpHandler interface {
	GetAll(c *echo.Context) error
	Create(c *echo.Context) error
}

type ItemUsecase interface {
	ItemList(filter *RequestItemFilter) (*ResponseItemList, error)
	CreateItem(req *RequestItemCreate) (*ItemModel, error)
}

type ItemRepository interface {
	Listing(filter *RequestItemFilter) ([]*ItemEntity, int, error)
	Create(item *ItemEntity) (*ItemEntity, error)
	Edit(item *ItemEntity) (*ItemEntity, error)
	Archive(itemId int) error
}
