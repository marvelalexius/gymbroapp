package auth

import (
	"encoding/json"

	"github.com/marvelalexius/gymbroapp/config"
	"github.com/marvelalexius/gymbroapp/internal/controller/request"
	"github.com/marvelalexius/gymbroapp/internal/controller/response"
	"github.com/marvelalexius/gymbroapp/internal/model"
	"github.com/marvelalexius/gymbroapp/internal/repository"
	"github.com/marvelalexius/gymbroapp/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	cfg      *config.Config
	userRepo repository.IUserRepo
}

func NewAuthService(cfg *config.Config, userRepo repository.IUserRepo) *AuthService {
	return &AuthService{cfg: cfg, userRepo: userRepo}
}

func (a *AuthService) Login(req request.LoginRequest) (*model.User, *utils.TokenHeader, error) {
	var authResponse *response.AuthResponse

	user, err := a.userRepo.FindByEmail(req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, gorm.ErrRecordNotFound
		}

		return nil, nil, err
	}

	marshaledUser, _ := json.Marshal(user)
	err = json.Unmarshal(marshaledUser, &authResponse)
	if err != nil {
		return nil, nil, err
	}

	err = a.verifyPassword(user, req.Password)
	if err != nil {
		return nil, nil, err
	}

	tokenHeader, err := a.generateAuthTokens(authResponse)
	if err != nil {
		return nil, nil, err
	}

	return user, tokenHeader, nil
}

func (a *AuthService) verifyPassword(u *model.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return bcrypt.ErrMismatchedHashAndPassword
		default:
			return err
		}
	}

	return err
}

func (a *AuthService) generateAuthTokens(user *response.AuthResponse) (*utils.TokenHeader, error) {
	refreshToken, err := utils.GenerateToken(user, a.cfg.App.RefreshTokenLifespan, a.cfg.App.TokenLifespanDuration, a.cfg.App.Secret)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user, a.cfg.App.TokenLifespan, a.cfg.App.TokenLifespanDuration, a.cfg.App.Secret)
	if err != nil {
		return nil, err
	}

	tokenHeader := utils.TokenHeader{
		AuthToken:           token.Token,
		AuthTokenExpires:    token.Expires,
		RefreshToken:        refreshToken.Token,
		RefreshTokenExpires: refreshToken.Expires,
	}

	return &tokenHeader, err
}
