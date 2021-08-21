package main

import (
	"go-fancy-todo/config"
	"go-fancy-todo/routes"
	"os"
)

func main() {
	e := routes.Init()
	port := os.Getenv("PORT")	

	config.NewDB()

	e.Logger.Fatal(e.Start(":"+port))
}

