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
	echoMiddleware "github.com/labstack/echo/v5/middleware"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/internal/adapter"
	"github.com/karan-khu/go-online-shop/internal/app/auth"
	"github.com/karan-khu/go-online-shop/internal/app/balance"
	"github.com/karan-khu/go-online-shop/internal/app/inventory"
	"github.com/karan-khu/go-online-shop/internal/app/item"
	"github.com/karan-khu/go-online-shop/internal/app/purchase"
	"github.com/karan-khu/go-online-shop/internal/app/user"
	"github.com/karan-khu/go-online-shop/internal/echovalidator"
	"github.com/karan-khu/go-online-shop/internal/middleware"
	"github.com/karan-khu/go-online-shop/internal/upload"
)

type Server struct {
	*echo.Echo
	imageBuilder upload.ImageBuilder
	conf         *config.Config
}

func NewServer() *Server {
	conf := config.NewConfig()
	app := echo.New()

	level := slog.LevelDebug
	if conf.Env.GO_ENV == "production" {
		level = slog.LevelError
	}
	app.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	app.Validator = echovalidator.NewValidator()

	app.Static("/uploads", "./uploads")

	imageBuilder := upload.NewImageBuilder(conf.Env.APP_HOST, "./uploads")

	return &Server{
		Echo:         app,
		conf:         conf,
		imageBuilder: imageBuilder,
	}
}

func main() {
	server := NewServer()

	// Register middleware
	server.Use(echoMiddleware.Recover())
	server.Use(middleware.RequestLogger(server.conf.Env))
	server.Use(middleware.CorsMiddleware(server.conf.Env))
	server.Use(middleware.RateLimit(100))
	server.Use(middleware.RequestTimeout(10 * time.Second))

	// Init endpoint
	server.GET("/api/v1/health", server.healthCheck)
	server.POST("/api/v1/upload/image", server.uploadImage)
	server.GET("*", server.notFound)

	// Register package dependency injection
	userRepo := user.NewUserRepository(server.Logger, server.conf)
	userCreator := adapter.NewAuthUserAdapter(userRepo)
	authGoogleUsecase := auth.NewAuthGoogleUsecase(userCreator)
	authGoogleHandler := auth.NewAuthGoogleHandler(server.Logger, server.conf, authGoogleUsecase)
	balanceRepo := balance.NewBalanceRepository(server.Logger, server.conf)
	balanceUsecase := balance.NewBalanceUsecase(balanceRepo)
	balanceHttpHandler := balance.NewBalanceHttpHandler(server.Logger, balanceUsecase)
	authMiddleware := middleware.NewAuthorizationMiddleware(server.Logger, server.conf, authGoogleUsecase)
	inventoryRepo := inventory.NewInventoryRepository(server.Logger, server.conf)
	inventoryUsecase := inventory.NewInventoryUsecase(inventoryRepo)
	inventoryHttpHandler := inventory.NewInventoryHttpHandler(server.Logger, inventoryUsecase)
	itemRepo := item.NewItemRepository(server.Logger, server.conf)
	purchaseRepo := purchase.NewPurchaseRepository(server.Logger, server.conf)
	purchaseUsecase := purchase.NewPurchaseUsecase(purchaseRepo)
	itemUsecase := item.NewItemUsecase(server.Logger, itemRepo, server.imageBuilder)
	itemHttpHandler := item.NewItemHttpHandler(itemUsecase)
	purchaseHttpHandler := purchase.NewPurchaseHttpHandler(server.Logger, purchaseUsecase)

	// Router registering
	auth.RegisterRoutes(server.Echo, authGoogleHandler)
	balance.RegisterBalanceRoutes(server.Echo, balanceHttpHandler, authMiddleware)
	inventory.RegisterRoutes(server.Echo, inventoryHttpHandler, authMiddleware)
	item.RegisterRoutes(server.Echo, itemHttpHandler, authMiddleware)
	purchase.RegisterRoutes(server.Echo, purchaseHttpHandler, authMiddleware)

	// Start server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	sc := echo.StartConfig{
		Address:         ":" + server.conf.Env.APP_PORT,
		GracefulTimeout: 10 * time.Second,
	}
	if err := sc.Start(ctx, server.Echo); err != nil {
		log.Fatal("Error starting server", "error", err)
	}
	log.Println("Server shutdown gracefully")
}

func (s *Server) healthCheck(pctx *echo.Context) error {
	return pctx.JSON(http.StatusOK, map[string]string{"message": "Server is running!"})
}

func (s *Server) notFound(pctx *echo.Context) error {
	return pctx.JSON(http.StatusNotFound, map[string]string{"error": "Not Found"})
}

func (s *Server) uploadImage(pctx *echo.Context) error {
	path, err := s.imageBuilder.SaveImage(pctx)
	if err != nil {
		return pctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return pctx.JSON(http.StatusOK, map[string]string{"url": path, "fullPath": s.imageBuilder.Build(path)})
}
