package middlewares

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var signingKey = []byte("privatekey")
var UserId float64

var Authentication = middleware.JWTWithConfig(middleware.JWTConfig{
	TokenLookup: "header:" + "access_token",
	// Claims:      jwt.MapClaims{},
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

// TODO get UserId from header

// func Decode(token string) {
// 	claims := jwt.MapClaims{}
// 	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("<YOUR VERIFICATION KEY>"), nil
// 	})

// 	UserId = claims["id"]
// }
