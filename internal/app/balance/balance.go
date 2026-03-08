package balance

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/internal/middleware"
)

func RegisterBalanceRoutes(app *echo.Echo, httpHandler BalanceHttpHandler, authMiddleware middleware.AuthorizationMiddleware) {
	router := app.Group("/api/v1/balance")

	router.GET("/show", httpHandler.CoinShowBalance, authMiddleware.Authorizing)
	router.POST("/topup", httpHandler.CoinTopUp, authMiddleware.Authorizing)
}
