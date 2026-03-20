package item

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/internal/response"
)

type itemHttpHandlerImpl struct {
	itemUsecase ItemUsecase
}

func NewItemHttpHandler(itemUsecase ItemUsecase) ItemHttpHandler {
	return &itemHttpHandlerImpl{
		itemUsecase: itemUsecase,
	}
}

func (h *itemHttpHandlerImpl) Create(pctx *echo.Context) error {
	adminId, ok := pctx.Get("userId").(string)
	if !ok || adminId == "" {
		return response.Unauthorized(pctx, errors.New("Unauthorized"))
	}

	req := new(RequestItemCreate)
	if err := pctx.Bind(req); err != nil {
		pctx.Logger().Error("failed to bind request", "error", err)
		return response.BadRequest(pctx, err)
	} else if err := pctx.Validate(req); err != nil {
		return response.BadRequest(pctx, err)
	}

	item, err := h.itemUsecase.CreateItem(req, adminId)
	if err != nil {
		pctx.Logger().Error("failed to create item", "error", err)
		return response.InternalError(pctx, err)
	}

	return response.Success(pctx, "Item created successfully", item)
}

func (h *itemHttpHandlerImpl) Edit(pctx *echo.Context) error {
	req := new(RequestItemEdit)
	if err := pctx.Bind(req); err != nil {
		pctx.Logger().Error("failed to bind request", "error", err)
		return response.BadRequest(pctx, err)
	} else if err := pctx.Validate(req); err != nil {
		return response.BadRequest(pctx, err)
	}

	item, err := h.itemUsecase.EditItem(req)
	if err != nil {
		pctx.Logger().Error("failed to edit item", "error", err)
		return response.InternalError(pctx, err)
	}

	return response.Success(pctx, "Item edited successfully", item)
}

func (h *itemHttpHandlerImpl) Delete(pctx *echo.Context) error {
	p_itemId := pctx.Param("item_id")
	if p_itemId == "" {
		return response.BadRequest(pctx, errors.New("item ID is required"))
	}
	itemId, ConvertErr := strconv.Atoi(p_itemId)
	if ConvertErr != nil {
		return response.BadRequest(pctx, ConvertErr)
	}

	if err := h.itemUsecase.DeleteItem(itemId); err != nil {
		pctx.Logger().Error("failed to delete item", "error", err)
		return response.BadRequest(pctx, err)
	}

	return response.Success(pctx, "Item deleted successfully", nil)
}
