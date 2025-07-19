package app

import (
	"backend_golang/config"
	"backend_golang/database"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type App struct {
	cfg    config.Config
	logger *logrus.Logger
}

func New(cfg config.Config, logger *logrus.Logger) *App {
	return &App{
		cfg:    cfg,
		logger: logger,
	}
}

func (a *App) initRoutes(r *gin.Engine, db *sql.DB) {
	// Create Transactor
	api := r.Group("/api")

	// Create Product Layer
	// Create Product Routes
	product := api.Group("/product")
	{
		product.GET("")
	}
}

func (a *App) serve(r *gin.Engine) {

	srv := &http.Server{
		Addr:    ":" + a.cfg.ServerPort,
		Handler: r.Handler(),
	}

	go func() {
		a.logger.Info("Server Starting...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.
				WithField("port", srv.Addr).
				WithField("error", err).
				Fatal("Unable to start the server ")
		}
	}()

}

func (a *App) Run() {
	db, err := database.ConnectDB(a.cfg)
	if err != nil {
		a.logger.WithField("error", err).Fatal("Unable to connect to the database")
	}
	defer db.Close()

	r := gin.New()
	r.Use(gin.Recovery())

	a.initRoutes(r, db)

	a.serve(r)
}
