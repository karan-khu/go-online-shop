package inventory

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/internal/middleware"
)

func RegisterRoutes(app *echo.Echo, httpHandler InventoryHttpHandler, authMiddleware middleware.AuthorizationMiddleware) {
	router := app.Group("/api/v1/inventory")

	router.GET("", httpHandler.Listing, authMiddleware.Authorizing)
}
