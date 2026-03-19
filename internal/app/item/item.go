package item

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/internal/middleware"
)

func RegisterRoutes(app *echo.Echo, httpHandler ItemHttpHandler, authMiddleware middleware.AuthorizationMiddleware) {
	v1 := app.Group("/api/v1/items")

	v1.POST("/create", httpHandler.Create, authMiddleware.Authorizing)
	v1.PUT("/edit", httpHandler.Edit, authMiddleware.Authorizing)
	v1.DELETE("/delete/:item_id", httpHandler.Delete, authMiddleware.Authorizing)
}
