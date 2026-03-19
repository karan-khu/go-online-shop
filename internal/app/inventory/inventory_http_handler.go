package inventory

import (
	"log/slog"

	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/internal/response"
)

type inventoryHttpHandlerImpl struct {
	logger           *slog.Logger
	inventoryUsecase InventoryUsecase
}

func NewInventoryHttpHandler(logger *slog.Logger, inventoryUsecase InventoryUsecase) InventoryHttpHandler {
	return &inventoryHttpHandlerImpl{
		logger:           logger,
		inventoryUsecase: inventoryUsecase,
	}
}

func (h *inventoryHttpHandlerImpl) Listing(pctx *echo.Context) error {
	userId := pctx.Get("userId").(string)
	inventory, err := h.inventoryUsecase.Listing(userId)
	if err != nil {
		h.logger.Error("failed to get inventory list", "error", err)
		return response.InternalError(pctx, err)
	}
	return response.Success(pctx, "Inventory list fetched successfully", inventory)
}
