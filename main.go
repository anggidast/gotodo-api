package main

import (
	"go-fancy-todo/routes"
	"os"
)

func main() {
	e := routes.Init()
	port := os.Getenv("PORT")	
	e.Logger.Fatal(e.Start(":"+port))
}
