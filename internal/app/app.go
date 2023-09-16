package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/marvelalexius/gymbroapp/config"
	v1 "github.com/marvelalexius/gymbroapp/internal/controller/http/v1"
	"github.com/marvelalexius/gymbroapp/internal/di"
	"github.com/marvelalexius/gymbroapp/internal/model"
	"github.com/marvelalexius/gymbroapp/pkg/httpserver"
	"github.com/marvelalexius/gymbroapp/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run(cfg *config.Config) {
	l := logger.NewLogger(cfg.Log.Level)
	db, err := gorm.Open(postgres.Open(cfg.PG.GetDbConnectionUrl()))

	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres: %v", err))
	}

	err = db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Migrate - postgres: %v", err))
	}

	di := di.NewDependencyInjection(db, cfg)

	handler := gin.New()
	v1.NewRouter(handler, db, cfg, di)
	httpServer := httpserver.NewServer(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app run: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("%w", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("%w", err))
	}
}
