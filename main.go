package main

import "go-fancy-todo/routes"

func main() {
	e := routes.Init()
	
	e.Logger.Fatal(e.Start(":8080"))
}
