package main

import (
	"go-fancy-todo/config"
	"go-fancy-todo/models"
	"go-fancy-todo/routes"
	"os"
)

func main() {
	e := routes.Init()
	port := os.Getenv("PORT")	
	e.Logger.Fatal(e.Start(":"+port))

	Todo := models.Todo{}
	User := models.User{}

	db := config.NewDB()

	db.Migrator().CreateTable(&User)
	db.Migrator().CreateTable(&Todo)
}

