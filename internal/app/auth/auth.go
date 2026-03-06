package auth

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/config"
)

func RegisterRoutes(app *echo.Echo, conf *config.Config, authUsecase AuthGoogleUsecase) {
	router := app.Group("/api/v1/auth")

	authGoogleHandler := NewAuthGoogleHandler(app.Logger, conf, authUsecase)

	router.GET("/google/login", authGoogleHandler.GoogleLogin)
	router.GET("/google/callback", authGoogleHandler.GoogleLoginCallBack)
	router.POST("/google/logout", authGoogleHandler.Logout)
}
