package models

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	User struct {
		ID        int       `json:"id"`
		Email     string    `json:"email" validate:"required,email"`
		Password  string    `gorm:"type:varchar(255)" json:"password" validate:"required"`
		CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
		UpdatedAt time.Time `json:"updatedAt"  gorm:"column:updatedAt"`
	}

	CustomValidator struct {
		Validator *validator.Validate
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

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
