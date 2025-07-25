package app

import (
	"backend_golang/config"
	"backend_golang/database"
	"backend_golang/internal/cache"
	"backend_golang/internal/handler"
	"backend_golang/internal/repository"
	"backend_golang/internal/usecase"
	"backend_golang/middleware"
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type App struct {
	cfg config.Config
	log *logrus.Logger
}

func New(cfg config.Config, logger *logrus.Logger) *App {
	return &App{
		cfg: cfg,
		log: logger,
	}
}

func (a *App) initRoutes(r *gin.Engine, db *sql.DB, redis *redis.Client) {
	api := r.Group("/api")

	// Create Product Layer
	productRepo := repository.NewProductRepo(db)
	productCache := cache.NewProductCache(redis)
	productUsecase := usecase.NewProductUsecase(productRepo, productCache)
	productHandler := handler.NewProductHandler(productUsecase)
	// Create Product Routes
	product := api.Group("/product")
	{
		product.GET("", productHandler.ListAllProduct)
		product.POST("", productHandler.AddProduct)
	}
}

func (a *App) serve(r *gin.Engine) {

	srv := &http.Server{
		Addr:    ":" + a.cfg.ServerPort,
		Handler: r,
	}

	go func() {
		a.log.Info("Server Starting...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.log.
				WithField("port", srv.Addr).
				WithField("error", err).
				Fatal("Unable to start the server ")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.log.Warnf("Server shutting down in %d seconds", a.cfg.ShutdownTime)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(a.cfg.ShutdownTime)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		a.log.WithField("error", err).Warn("Server failed to shutdown")
	}
	<-ctx.Done()

	a.log.Warn("Server exiting...")
}

func (a *App) Run() {
	db, err := database.ConnectDB(a.cfg)
	if err != nil {
		a.log.WithField("error", err).Fatal("Unable to connect to the database")
	}
	defer db.Close()

	redisClient, err := database.ConnectRedis(a.cfg)
	if err != nil {
		a.log.WithField("error", err).Fatal("Unable to connect to Redis")
	}
	defer redisClient.Close()

	r := gin.New()
	r.ContextWithFallback = true
	r.Use(gin.Recovery())
	r.Use(middleware.Logger(a.log))
	r.Use(middleware.ErrorMiddleware(a.log))

	a.initRoutes(r, db, redisClient)
	a.serve(r)
}
