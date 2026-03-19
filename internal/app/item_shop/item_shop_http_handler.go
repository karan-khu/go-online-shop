package itemshop

import (
	"errors"

	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/pkg/res"
)

type itemShopHttpHandler struct {
	itemShopUsecase ItemShopUsecase
}

func NewItemShopHttpHandler(itemShopUsecase ItemShopUsecase) ItemShopHttpHandler {
	return &itemShopHttpHandler{
		itemShopUsecase: itemShopUsecase,
	}
}

func (h *itemShopHttpHandler) GetAll(pctx *echo.Context) error {
	req := &RequestItemFilter{
		Page:  1,
		Limit: 10,
	}
	if err := pctx.Bind(req); err != nil {
		pctx.Logger().Error("failed to bind request", "error", err)
		return res.BadRequest(pctx, err)
	}

	list, err := h.itemShopUsecase.ItemList(req)
	if err != nil {
		pctx.Logger().Error("failed to get item list", "error", err)
		return res.InternalError(pctx, err)
	}

	return res.Success(pctx, "Item list fetched successfully", list)
}

func (h *itemShopHttpHandler) Buying(pctx *echo.Context) error {
	userId := pctx.Get("userId").(string)
	if userId == "" {
		return res.Unauthorized(pctx, errors.New("Unauthorized"))
	}

	req := new(RequestItemBuying)
	if err := pctx.Bind(req); err != nil {
		pctx.Logger().Error("failed to bind request", "error", err)
		return res.BadRequest(pctx, err)
	} else if err := pctx.Validate(req); err != nil {
		return res.BadRequest(pctx, err)
	}

	req.UserId = userId
	if err := h.itemShopUsecase.Buying(req); err != nil {
		return res.BadRequest(pctx, err)
	}

	return res.Success(pctx, "Buying item successfully", nil)
}

func (h *itemShopHttpHandler) Selling(pctx *echo.Context) error {
	userId := pctx.Get("userId").(string)
	if userId == "" {
		return res.Unauthorized(pctx, errors.New("Unauthorized"))
	}

	req := new(RequestItemSelling)
	if err := pctx.Bind(req); err != nil {
		pctx.Logger().Error("failed to bind request", "error", err)
		return res.BadRequest(pctx, err)
	} else if err := pctx.Validate(req); err != nil {
		return res.BadRequest(pctx, err)
	}

	req.UserId = userId
	if err := h.itemShopUsecase.Selling(req); err != nil {
		return res.BadRequest(pctx, err)
	}

	return res.Success(pctx, "Selling item successfully", nil)
}
