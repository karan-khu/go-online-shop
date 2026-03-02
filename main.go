package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/karan-khu/go-online-shop/config"
)

func main() {
	conf := config.NewConfig()

	app := echo.New()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: conf.Env.APP_LOG_LEVEL,
	}))
	app.Logger = logger

	app.Use(middleware.RequestLogger())

	app.GET("/api/v1/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Server is running!"})
	})

	port := fmt.Sprintf(":%s", conf.Env.APP_PORT)
	if err := app.Start(port); err != nil {
		log.Fatal("Error starting server", "error", err)
	}
}
