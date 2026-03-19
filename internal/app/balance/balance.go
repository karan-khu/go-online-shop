package balance

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/internal/middleware"
)

func RegisterBalanceRoutes(app *echo.Echo, httpHandler BalanceHttpHandler, authMiddleware middleware.AuthorizationMiddleware) {
	v1 := app.Group("/api/v1/balance")

	v1.GET("/show", httpHandler.CoinShowBalance, authMiddleware.Authorizing)
	v1.POST("/topup", httpHandler.CoinTopUp, authMiddleware.Authorizing)
}
