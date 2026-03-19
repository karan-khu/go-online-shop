package itemshop

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/internal/middleware"
)

func RegisterRoutes(app *echo.Echo, httpHandler ItemShopHttpHandler, authMiddleware middleware.AuthorizationMiddleware) {
	v1 := app.Group("/api/v1/item-shop")

	v1.GET("", httpHandler.GetAll, authMiddleware.Authorizing)
	v1.POST("/buy", httpHandler.Buying, authMiddleware.Authorizing)
	v1.POST("/sell", httpHandler.Selling, authMiddleware.Authorizing)
}
