package item

import "github.com/labstack/echo/v5"

type ItemHttpHandler interface {
	GetAll(c *echo.Context) error
}

type ItemUsecase interface {
	ItemList() ([]*ItemModel, error)
}

type ItemRepository interface {
	Listing() ([]*ItemEntity, error)
}
