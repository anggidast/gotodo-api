package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

var Authentication = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("privatekey"),
})

// TODO get UserId from header
var UserId interface{}

func Decode(token string) {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("<YOUR VERIFICATION KEY>"), nil
	})

	UserId = claims["id"]
}
