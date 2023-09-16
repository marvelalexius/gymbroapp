package response

import "time"

type (
	AuthResponse struct {
		ID                     uint      `json:"id"`
		FullName               string    `json:"fullName" example:"user name"`
		Email                  string    `json:"email" example:"email@email.com"`
		Password               string    `json:"-" example:"password123"`
		Weight                 *uint     `json:"weight,omitempty"`
		Height                 *uint     `json:"height,omitempty"`
		Age                    uint      `json:"age,omitempty" gorm:"-"`
		Overview               *string   `json:"overview,omitempty"`
		RefreshToken           string    `json:"-"`
		RefreshTokenExpiration string    `json:"-"`
		CreatedAt              time.Time `json:"createdAt,omitempty" example:"2023-01-01T15:01:00+00:00"`
		UpdatedAt              time.Time `json:"updatedAt,omitempty" example:"2023-02-11T15:01:00+00:00"`
		Token                  string    `json:"token,omitempty"`
		Expires                time.Time `json:"expires,omitempty"`
	}
)
