package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	md "github.com/labstack/echo/v5/middleware"
)

func RateLimit(limit int) echo.MiddlewareFunc {
	return md.RateLimiterWithConfig(md.RateLimiterConfig{
		Skipper: md.DefaultSkipper,
		Store: md.NewRateLimiterMemoryStoreWithConfig(md.RateLimiterMemoryStoreConfig{
			Burst:     limit,
			ExpiresIn: time.Minute,
		}),
		ErrorHandler: func(c *echo.Context, err error) error {
			return c.JSON(http.StatusTooManyRequests, map[string]string{"error": "Rate limit exceeded."})
		},
	})
}
