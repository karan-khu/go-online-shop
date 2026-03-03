package item

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/pkg/res"
)

type itemHttpHandlerImpl struct {
	itemUsecase ItemUsecase
}

func NewItemHttpHandler(itemUsecase ItemUsecase) ItemHttpHandler {
	return &itemHttpHandlerImpl{
		itemUsecase: itemUsecase,
	}
}

func (h *itemHttpHandlerImpl) GetAll(c *echo.Context) error {
	list, err := h.itemUsecase.ItemList()
	if err != nil {
		return res.InternalError(c, err)
	}

	return res.Success(c, "Item list fetched successfully", list)
}
