package service

import (
	"github.com/marvelalexius/gymbroapp/internal/controller/request"
	"github.com/marvelalexius/gymbroapp/internal/model"
	"github.com/marvelalexius/gymbroapp/pkg/utils"
)

type (
	IAuthService interface {
		Login(req request.LoginRequest) (*model.User, *utils.TokenHeader, error)
	}
)
