package controllers

import (
	"fmt"
	"go-fancy-todo/config"
	"go-fancy-todo/helpers"
	"go-fancy-todo/models"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
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
	c.Request().Header.Set(echo.HeaderAccessControlAllowOrigin, "*")
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	db := config.NewDB()
	user := models.User{}

	if err = db.First(&user, "email = ?", req.Email).Error; err != nil {
		return echo.NewHTTPError(http.StatusForbidden, map[string]string{
			"message": "Email is not registered",
		})
	} else {
		if err = helpers.CompareHash(user.Password, req.Password); err != nil {
			return echo.NewHTTPError(http.StatusForbidden, map[string]string{
				"message": "Wrong password",
			})
		}
	}

	// ? JWT
	// * JWT init
	token := jwt.New(jwt.SigningMethodHS256)

	// * define payload
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// * generate encoded token
	t, err := token.SignedString([]byte("privatekey"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	// * send it as response

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "succeed",
		"user_id":      user.ID,
		"access_token": t,
	})
}

func Register(c echo.Context) (err error) {
	c.Request().Header.Set(echo.HeaderAccessControlAllowOrigin, "*")
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

	req := new(models.User)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db := config.NewDB()
	user := models.User{}

	if err = c.Validate(req); err != nil {
		return err
	}

	if err = db.First(&user, "email = ?", req.Email).Error; err == nil {
		return echo.NewHTTPError(http.StatusForbidden, map[string]string{
			"message": fmt.Sprintf("Email %v is already registered", req.Email),
		})
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
