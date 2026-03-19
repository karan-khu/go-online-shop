package auth

import (
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(app *echo.Echo, googleHandler AuthGoogleHandler) {
	v1 := app.Group("/api/v1/auth")

	v1.GET("/google/login", googleHandler.GoogleLogin)
	v1.GET("/google/callback", googleHandler.GoogleLoginCallBack)
	v1.POST("/google/logout", googleHandler.Logout)
}
