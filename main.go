package main

import (
	"go-fancy-todo/config"
	"go-fancy-todo/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main(c echo.Context) {
	e := routes.Init(c)
	port := os.Getenv("PORT")

	config.NewDB()

	e.Logger.Fatal(e.Start(":" + port))
}
