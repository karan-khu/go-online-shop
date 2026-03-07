package balance

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/internal/middleware"
)

func RegisterBalanceRoutes(app *echo.Echo, conf *config.Config, authMiddleware middleware.AuthorizationMiddleware) {

	router := app.Group("/api/v1/balance")

	balanceRepo := NewBalanceRepository(app.Logger, conf)
	balanceUsecase := NewBalanceUsecase(balanceRepo)
	balanceHttpHandler := NewBalanceHttpHandler(app.Logger, balanceUsecase)

	router.GET("/show", balanceHttpHandler.CoinShowBalance, authMiddleware.Authorizing)
	router.POST("/topup", balanceHttpHandler.CoinTopUp, authMiddleware.Authorizing)
}
