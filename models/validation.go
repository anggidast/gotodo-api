package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
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
		case "Key: 'Todo.Title' Error:Field validation for 'Title' failed on the 'required' tag Key: 'Todo.Due_date' Error:Field validation for 'Due_date' failed on the 'required' tag":
			return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
				"message": "Todo title and due date cannot be empty/null",
			})
		case "Key: 'Todo.Due_date' Error:Field validation for 'Due_date' failed on the 'required' tag":
			return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
				"message": "Todo title and due date cannot be empty/null",
			})
		case "Key: 'Todo.Title' Error:Field validation for 'Title' failed on the 'required' tag":
			return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
				"message": "Todo title and due date cannot be empty/null",
			})
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
