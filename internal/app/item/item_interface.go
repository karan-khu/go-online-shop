package item

import "github.com/labstack/echo/v5"

type ItemHttpHandler interface {
	GetAll(c *echo.Context) error
}

type ItemUsecase interface {
	ItemList(filter *RequestItemFilter) (*ResponseItemList, error)
}

type ItemRepository interface {
	Listing(filter *RequestItemFilter) ([]*ItemEntity, int, error)
}
