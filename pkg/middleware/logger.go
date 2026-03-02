package middleware

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/labstack/echo/v5"
	echoMiddleware "github.com/labstack/echo/v5/middleware"

	"github.com/karan-khu/go-online-shop/config"
)

var (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
)

var methodColors = map[string]string{
	"GET":    colorBlue,
	"POST":   colorCyan,
	"PUT":    colorYellow,
	"PATCH":  colorYellow,
	"DELETE": colorRed,
}

func SetLogger(app *echo.Echo, env *config.Env) echo.MiddlewareFunc {
	level := slog.LevelDebug
	if env.GO_ENV == "production" {
		level = slog.LevelError
	}

	app.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	return echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogLatency:  true,
		HandleError: true,
		LogValuesFunc: func(c *echo.Context, v echoMiddleware.RequestLoggerValues) error {
			timestamp := time.Now().Format("2006/01/02 - 03:04:05")
			mc := methodColors[v.Method]
			sc := func() string {
				if v.Status >= 200 && v.Status < 300 {
					return colorGreen
				} else if v.Status >= 300 && v.Status < 400 {
					return colorBlue
				}
				return colorRed
			}()
			fmt.Printf("[ECHO] %s | %s%d%s | %s | %s%s%s %s\n",
				timestamp,
				sc, v.Status, colorReset,
				v.Latency,
				mc, v.Method, colorReset,
				v.URI,
			)
			return nil
		},
	})
}
