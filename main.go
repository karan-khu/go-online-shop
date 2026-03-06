package main

import (
	"context"
	"fmt"
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
	"github.com/karan-khu/go-online-shop/internal/adapter"
	"github.com/karan-khu/go-online-shop/internal/app/auth"
	"github.com/karan-khu/go-online-shop/internal/app/item"
	"github.com/karan-khu/go-online-shop/internal/app/user"
	md "github.com/karan-khu/go-online-shop/internal/middleware"
	"github.com/karan-khu/go-online-shop/pkg/upload"
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
	app.Use(md.RequestLogger(conf.Env))
	app.Use(md.CorsMiddleware(conf.Env))

	app.Static("/uploads", "./uploads")
	app.POST("/api/v1/upload/image", func(c *echo.Context) error {
		path, err := upload.SaveImage(c, "./uploads")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]string{"url": path, "fullPath": fmt.Sprintf("%s/%s", conf.Env.APP_HOST, path)})
	})
	app.GET("/api/v1/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Server is running!"})
	})

	userRepo := user.NewUserRepository(app.Logger, conf)
	userCreator := adapter.NewAuthUserAdapter(userRepo)
	authUsecase := auth.NewAuthGoogleUsecase(userCreator)
	authMiddleware := md.NewAuthorizationMiddleware(app.Logger, conf, authUsecase)

	auth.RegisterRoutes(app, conf, authUsecase)
	item.RegisterRoutes(app, conf, authMiddleware)

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
