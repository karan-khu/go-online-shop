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
	req := &RequestItemFilter{
		Page:  1,
		Limit: 10,
	}
	if err := c.Bind(req); err != nil {
		c.Logger().Error("failed to bind request", "error", err)
		return res.BadRequest(c, "Invalid request", err)
	}

	list, err := h.itemUsecase.ItemList(req)
	if err != nil {
		c.Logger().Error("failed to get item list", "error", err)
		return res.InternalError(c, err)
	}

	return res.Success(c, "Item list fetched successfully", list)
}
