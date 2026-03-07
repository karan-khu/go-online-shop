package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	md "github.com/labstack/echo/v5/middleware"
)

func RequestTimeout(timeOut time.Duration) echo.MiddlewareFunc {
	return md.ContextTimeoutWithConfig(md.ContextTimeoutConfig{
		Skipper: md.DefaultSkipper,
		ErrorHandler: func(c *echo.Context, err error) error {
			return c.JSON(http.StatusRequestTimeout, map[string]string{"error": "Request Timeout."})
		},
		Timeout: timeOut * time.Second,
	})
}
