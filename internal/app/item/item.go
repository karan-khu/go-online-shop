package item

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/internal/middleware"
)

func RegisterRoutes(app *echo.Echo, conf *config.Config, authMiddleware middleware.AuthorizationMiddleware) {
	router := app.Group("/api/v1/items")

	itemRepo := NewItemRepository(app.Logger, conf)
	itemUsecase := NewItemUsecase(app.Logger, itemRepo, conf)
	itemHttpHandler := NewItemHttpHandler(itemUsecase)

	router.GET("", itemHttpHandler.GetAll, authMiddleware.Authorizing)
	router.POST("/create", itemHttpHandler.Create)
	router.PUT("/edit", itemHttpHandler.Edit)
	router.DELETE("/delete/:item_id", itemHttpHandler.Delete)
}
