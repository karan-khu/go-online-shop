package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/cors"
	"github.com/karan-khu/go-online-shop/pkg/logger"
	"github.com/karan-khu/go-online-shop/pkg/validator"
)

func main() {
	conf := config.NewConfig()

	app := echo.New()
	level := slog.LevelDebug
	if conf.Env.GO_ENV == "production" {
		level = slog.LevelError
	}
	app.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	app.Validator = validator.NewValidator()

	app.Use(middleware.Recover())
	app.Use(logger.RequestLogger(conf.Env))
	app.Use(cors.CorsMiddleware(conf.Env))

	app.GET("/api/v1/health", func(c *echo.Context) error {
		type HealthResponse struct {
			Message string `json:"message" validate:"required"`
		}
		req := new(HealthResponse)
		if err := validator.ValidateSchema(c, req); err != nil {
			return err
		}
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
