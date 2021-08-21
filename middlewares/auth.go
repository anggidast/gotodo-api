package middlewares

import (
	"errors"
	"fmt"
	"go-fancy-todo/config"
	"go-fancy-todo/models"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

var signingKey = []byte("privatekey")
var UserId float64

var Authentication = middleware.JWTWithConfig(middleware.JWTConfig{
	TokenLookup: "header:" + "access_token",
	ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
		keyFunc := func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != "HS256" {
				return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
			}
			return signingKey, nil
		}

		// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
		token, err := jwt.Parse(auth, keyFunc)
		if err != nil {
			return nil, err
		}
		if !token.Valid {
			return nil, errors.New("invalid token")
		}
		UserId = token.Claims.(jwt.MapClaims)["id"].(float64)
		return token, nil
	},
})

func Authorization(id string, c echo.Context) (todo models.Todo, userId string, db *gorm.DB, err error) {
	db = config.Database()
	todo = models.Todo{}

	userId = strconv.Itoa(int(UserId))

	if err = db.First(&todo, "user_id = ? AND id = ?", userId, id).Error; err != nil {
		return todo, userId, db, echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "Todo not found",
		})
	}

	return todo, userId, db, nil
}
