package item

import (
	"github.com/labstack/echo/v5"
)

type ItemHttpHandler interface {
	Create(c *echo.Context) error
	Edit(c *echo.Context) error
	Delete(c *echo.Context) error
}

type ItemUsecase interface {
	CreateItem(req *RequestItemCreate, adminId string) (*ItemEntity, error)
	EditItem(req *RequestItemEdit) (*ItemEntity, error)
	DeleteItem(itemId int) error
}

type ItemRepository interface {
	FindById(itemId int) (*ItemEntity, error)
	FindExists(itemId int) bool
	Create(item *ItemEntity) (*ItemEntity, error)
	Edit(item *ItemEntity) (*ItemEntity, error)
	Archive(itemId int) error
}
