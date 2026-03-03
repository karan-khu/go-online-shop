package item

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
)

type itemHttpHandlerImpl struct {
	logger      *slog.Logger
	itemUsecase ItemUsecase
}

func NewItemHttpHandler(logger *slog.Logger, itemUsecase ItemUsecase) ItemHttpHandler {
	return &itemHttpHandlerImpl{
		logger:      logger,
		itemUsecase: itemUsecase,
	}
}

func (h *itemHttpHandlerImpl) GetAll(c *echo.Context) error {
	list, err := h.itemUsecase.ItemList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Item list fetched successfully",
		"result":  list,
	})
}
