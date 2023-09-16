package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marvelalexius/gymbroapp/config"
	"github.com/marvelalexius/gymbroapp/internal/di"
	"github.com/marvelalexius/gymbroapp/internal/middleware"
	"gorm.io/gorm"
)

func NewRouter(handler *gin.Engine, db *gorm.DB, cfg *config.Config, di *di.DependencyInjection) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(middleware.CORS())

	handler.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "oks"})
	})

	h := handler.Group("api/v1")
	{
		newAuthRoutes(h, cfg, di.AuthService)
	}
}
