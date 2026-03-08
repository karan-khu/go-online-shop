package auth

import (
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(app *echo.Echo, googleHandler AuthGoogleHandler) {
	router := app.Group("/api/v1/auth")

	router.GET("/google/login", googleHandler.GoogleLogin)
	router.GET("/google/callback", googleHandler.GoogleLoginCallBack)
	router.POST("/google/logout", googleHandler.Logout)
}
