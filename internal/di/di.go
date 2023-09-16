package di

import (
	"github.com/marvelalexius/gymbroapp/config"
	userR "github.com/marvelalexius/gymbroapp/internal/repository/user"
	"github.com/marvelalexius/gymbroapp/internal/service/auth"
	"gorm.io/gorm"
)

type DependencyInjection struct {
	AuthService *auth.AuthService
}

func NewDependencyInjection(db *gorm.DB, cfg *config.Config) *DependencyInjection {
	userRepo := userR.NewUserRepo(db)

	authService := auth.NewAuthService(cfg, userRepo)

	return &DependencyInjection{
		AuthService: authService,
	}
}
