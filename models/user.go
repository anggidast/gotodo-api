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

	errorString struct {
		s string
	}
)

func (e *errorString) Error() string {
	return e.s
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code

		switch err.Error() {
		case "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag":
			return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
				"message": "Email cannot be empty/null",
			})
		case "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag":
			return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
				"message": "Email format is wrong",
			})
		case "Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag":
			return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
				"message": "Password cannot be empty/null",
			})
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
