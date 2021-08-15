package controllers

import (
	"go-fancy-todo/config"
	"go-fancy-todo/helpers"
	"go-fancy-todo/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) (err error) {
	db := config.NewDB()
	users := []models.User{}

	if err = db.Find(&users).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func Login(c echo.Context) (err error) {
	req := new(models.User)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	db := config.NewDB()
	user := models.User{}

	if err = db.First(&user, "email = ?", req.Email).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else {
		if err = helpers.CompareHash(user.Password, req.Password); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
	}

	response := models.UserResponse{
		Message: "succeed",
		Data:    models.UserInfo{ID: user.ID, Email: user.Email},
	}

	return c.JSON(http.StatusOK, response)
}

func Register(c echo.Context) (err error) {
	req := new(models.User)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	newUser := models.User{
		Email:     req.Email,
		Password:  helpers.Hash(req.Password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db := config.NewDB()
	db.Create(&newUser)
	response := models.UserResponse{
		Message: "created",
		Data:    models.UserInfo{ID: newUser.ID, Email: req.Email},
	}

	return c.JSON(http.StatusCreated, response)
}

func UpdateUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))

	db := config.NewDB()
	user := models.User{}

	if err = db.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	req := new(models.User)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user.Email = req.Email
	user.Password = req.Password
	user.UpdatedAt = time.Now()

	db.Save(&user)

	response := models.UserResponse{
		Message: "updated",
		Data:    models.UserInfo{ID: user.ID, Email: user.Email},
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))

	db := config.NewDB()
	user := models.User{}

	if err = db.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	db.Delete(&user, id)

	response := models.UserResponse{
		Message: "deleted",
		Data:    models.UserInfo{ID: user.ID, Email: user.Email},
	}

	return c.JSON(http.StatusOK, response)
}
