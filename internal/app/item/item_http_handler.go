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

func (h *itemHttpHandlerImpl) GetAll(pctx *echo.Context) error {
	req := &RequestItemFilter{
		Page:  1,
		Limit: 10,
	}
	if err := pctx.Bind(req); err != nil {
		pctx.Logger().Error("failed to bind request", "error", err)
		return res.BadRequest(pctx, "Invalid request", err)
	}

	list, err := h.itemUsecase.ItemList(req)
	if err != nil {
		pctx.Logger().Error("failed to get item list", "error", err)
		return res.InternalError(pctx, err)
	}

	return res.Success(pctx, "Item list fetched successfully", list)
}

func (h *itemHttpHandlerImpl) Create(pctx *echo.Context) error {
	req := new(RequestItemCreate)
	if err := pctx.Bind(req); err != nil {
		pctx.Logger().Error("failed to bind request", "error", err)
		return res.BadRequest(pctx, "Invalid request", err)
	}
	if err := pctx.Validate(req); err != nil {
		return res.BadRequest(pctx, "Invalid request", err)
	}

	item, err := h.itemUsecase.CreateItem(req)
	if err != nil {
		pctx.Logger().Error("failed to create item", "error", err)
		return res.InternalError(pctx, err)
	}

	return res.Success(pctx, "Item created successfully", item)
}
