package model

import "time"

type (
	User struct {
		ID                     uint   `gorm:"primary_key" json:"id"`
		FullName               string `json:"fullName" gorm:"not null" example:"user name"`
		Email                  string `json:"email" gorm:"not null;unique" example:"email@email.com"`
		Password               string `json:"-" gorm:"not null" example:"password123"`
		Weight                 *uint  `json:"weight,omitempty"`
		Height                 *uint  `json:"height,omitempty"`
		Age                    uint   `json:"age,omitempty" gorm:"-"`
		RefreshToken           string `json:"-"`
		RefreshTokenExpiration string `json:"-"`

		CreatedAt time.Time `json:"createdAt" example:"2023-01-01T15:01:00+00:00"`
		UpdatedAt time.Time `json:"updatedAt" example:"2023-02-11T15:01:00+00:00"`
	}
)
