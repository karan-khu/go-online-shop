package purchase

import (
	"errors"
	"log/slog"

	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/internal/response"
)

type purchaseHttpHandlerImpl struct {
	logger          *slog.Logger
	purchaseUsecase PurchaseUsecase
}

func NewPurchaseHttpHandler(logger *slog.Logger, purchaseUsecase PurchaseUsecase) PurchaseHttpHandler {
	return &purchaseHttpHandlerImpl{
		logger:          logger,
		purchaseUsecase: purchaseUsecase,
	}
}

func (h *purchaseHttpHandlerImpl) Listing(c *echo.Context) error {
	userId := c.Get("userId").(string)
	if userId == "" {
		return response.Unauthorized(c, errors.New("Unauthorized"))
	}

	purchase, err := h.purchaseUsecase.Listing(userId)
	if err != nil {
		h.logger.Error("failed to get purchase list", "error", err)
		return response.InternalError(c, err)
	}
	return response.Success(c, "Purchase list fetched successfully", purchase)
}
