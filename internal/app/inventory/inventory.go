package inventory

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/internal/middleware"
)

func RegisterRoutes(app *echo.Echo, conf *config.Config, authMiddleware middleware.AuthorizationMiddleware) {
	router := app.Group("/api/v1/inventory")

	inventoryRepo := NewInventoryRepository(app.Logger, conf)
	inventoryUsecase := NewInventoryUsecase(app.Logger, inventoryRepo)
	inventoryHttpHandler := NewInventoryHttpHandler(app.Logger, inventoryUsecase)

	router.GET("", inventoryHttpHandler.Listing, authMiddleware.Authorizing)
}
