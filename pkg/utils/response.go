package utils

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	TokenHeader struct {
		AuthToken           string
		AuthTokenExpires    time.Time
		RefreshToken        string
		RefreshTokenExpires time.Time
	}

	ValidationErrorMsg struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ErrorRes struct {
		Message string `json:"message"`
		Debug   error  `json:"debug,omitempty"`
		Errors  any    `json:"errors"`
	}

	SuccessRes struct {
		Message string      `json:"message"`
		Data    any         `json:"data,omitempty"`
		Header  TokenHeader `json:"-"`
	}
)

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " should be a valid email address"
	case "file":
		return fe.Field() + " should be a valid file"
	case "lte":
		return fe.Field() + " should be less than " + fe.Param()
	case "gte":
		return fe.Field() + " should be greater than " + fe.Param()
	case "eqfield":
		return fe.Field() + " should be equal to " + fe.Param()
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters long"
	case "max":
		return fe.Field() + " must be at most " + fe.Param() + " characters"
	default:
		return "Unknown error"
	}
}

func ValidationResponse(err error) []ValidationErrorMsg {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]ValidationErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ValidationErrorMsg{fe.Field(), getErrorMsg(fe)}
		}
		return out
	}

	return nil
}

func ErrorResponse(c *gin.Context, code int, res ErrorRes) {
	c.JSON(code, res)
}

func SuccessResponse(c *gin.Context, code int, res SuccessRes) {
	if res.Header != (TokenHeader{}) {
		c.Header("refresh-token", res.Header.RefreshToken)
		c.Header("refresh-token-expired", res.Header.RefreshTokenExpires.String())
		c.Header("Authorization", "Bearer "+res.Header.AuthToken)
		c.Header("expired-at", res.Header.AuthTokenExpires.String())
	}
	c.JSON(code, res)
}
