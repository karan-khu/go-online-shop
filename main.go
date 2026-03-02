package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/karan-khu/go-online-shop/config"
	md "github.com/karan-khu/go-online-shop/pkg/middleware"
)

func main() {
	conf := config.NewConfig()

	app := echo.New()

	app.Use(middleware.Recover())
	app.Use(md.SetLogger(app, conf.Env))
	app.Use(md.CorsMiddleware(conf.Env))

	app.GET("/api/v1/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Server is running!"})
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	sc := echo.StartConfig{
		Address:         ":" + conf.Env.APP_PORT,
		GracefulTimeout: 10 * time.Second,
	}
	if err := sc.Start(ctx, app); err != nil {
		log.Fatal("Error starting server", "error", err)
	}
	log.Println("Server shutdown gracefully")
}
