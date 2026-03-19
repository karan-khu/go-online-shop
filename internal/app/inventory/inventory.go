package inventory

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/internal/middleware"
)

func RegisterRoutes(app *echo.Echo, httpHandler InventoryHttpHandler, authMiddleware middleware.AuthorizationMiddleware) {
	v1 := app.Group("/api/v1/inventory")

	v1.GET("", httpHandler.Listing, authMiddleware.Authorizing)
}
