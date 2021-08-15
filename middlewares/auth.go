package middlewares

import (
	"github.com/labstack/echo/v4/middleware"
)

var Authentication = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("privatekey"),
})
