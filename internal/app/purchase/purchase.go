package purchase

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/internal/middleware"
)

func RegisterRoutes(app *echo.Echo, httpHandler PurchaseHttpHandler, authMiddleware middleware.AuthorizationMiddleware) {
	router := app.Group("/api/v1/purchase")

	router.GET("", httpHandler.Listing, authMiddleware.Authorizing)
}
