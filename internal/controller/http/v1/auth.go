package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marvelalexius/gymbroapp/config"
	"github.com/marvelalexius/gymbroapp/internal/controller/request"
	"github.com/marvelalexius/gymbroapp/internal/service"
	"github.com/marvelalexius/gymbroapp/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authRoutes struct {
	cfg *config.Config
	s   service.IAuthService
}

func newAuthRoutes(handler *gin.RouterGroup, cfg *config.Config, s service.IAuthService) {
	r := &authRoutes{s: s, cfg: cfg}

	h := handler.Group("auth")
	{
		h.POST("login", r.login)
	}
}

func (r *authRoutes) login(ctx *gin.Context) {
	var req request.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ve := utils.ValidationResponse(err)

		utils.ErrorResponse(ctx, http.StatusBadRequest, utils.ErrorRes{
			Message: "request not valid",
			Debug:   nil,
			Errors:  ve,
		})
		return
	}

	user, token, err := r.s.Login(req)
	if err != nil {
		statusCode := http.StatusBadRequest
		errors := err.Error()

		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
		} else if err == bcrypt.ErrMismatchedHashAndPassword {
			statusCode = http.StatusUnprocessableEntity
			errors = "password is incorrect"
		}

		utils.ErrorResponse(ctx, statusCode, utils.ErrorRes{
			Message: "Something went wrong",
			Debug:   err,
			Errors:  errors,
		})
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, utils.SuccessRes{
		Message: "Login Successful",
		Data:    user,
		Header:  *token,
	})
}
