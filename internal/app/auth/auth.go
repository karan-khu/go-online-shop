package auth

import (
	"github.com/labstack/echo/v5"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/internal/app/user"
)

func RegisterRoutes(app *echo.Echo, conf *config.Config) {
	router := app.Group("/api/v1/auth")

	userRepo := user.NewUserRepository(app.Logger, conf)

	authGoogleUsecase := NewAuthGoogleUsecase(userRepo)
	authGoogleHandler := NewAuthGoogleHandler(app.Logger, conf, authGoogleUsecase)

	router.GET("/google/login", authGoogleHandler.GoogleLogin)
}
