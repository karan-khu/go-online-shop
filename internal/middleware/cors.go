package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/karan-khu/go-online-shop/config"
)

func CorsMiddleware(env *config.Env) echo.MiddlewareFunc {
	whitelist := strings.Split(env.APP_CORS, ",")

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     whitelist,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	})
}
