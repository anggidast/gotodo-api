package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type (
	User struct {
		ID        int       `json:"id"`
		Email     string    `json:"email" validate:"required, email"`
		Password  string    `gorm:"type:varchar(255)" json:"password" validate:"required"`
		CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
		UpdatedAt time.Time `json:"updatedAt"  gorm:"column:updatedAt"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}

	UserInfo struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
	}

	UserResponse struct {
		Message string   `json:"message"`
		Data    UserInfo `json:"data"`
	}
)
