package item

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/config"
)

func RegisterRoutes(app *echo.Echo, conf *config.Config) {
	router := app.Group("/api/v1/items")

	itemRepo := NewItemRepository(app.Logger, conf)
	itemUsecase := NewItemUsecase(app.Logger, itemRepo, conf)
	itemHttpHandler := NewItemHttpHandler(itemUsecase)

	router.GET("", itemHttpHandler.GetAll)
	router.POST("/create", itemHttpHandler.Create)
	router.PUT("/edit", itemHttpHandler.Edit)
	router.DELETE("/delete/:item_id", itemHttpHandler.Delete)
}
