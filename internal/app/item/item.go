package item

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/internal/middleware"
)

func RegisterRoutes(app *echo.Echo, httpHandler ItemHttpHandler, authMiddleware middleware.AuthorizationMiddleware) {
	router := app.Group("/api/v1/items")

	router.GET("", httpHandler.GetAll, authMiddleware.Authorizing)
	router.POST("/create", httpHandler.Create, authMiddleware.Authorizing)
	router.PUT("/edit", httpHandler.Edit, authMiddleware.Authorizing)
	router.DELETE("/delete/:item_id", httpHandler.Delete, authMiddleware.Authorizing)
	router.POST("/buy", httpHandler.Buying, authMiddleware.Authorizing)
	router.POST("/sell", httpHandler.Selling, authMiddleware.Authorizing)
}
