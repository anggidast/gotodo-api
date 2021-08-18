package main

import (
	"go-fancy-todo/config"
	"go-fancy-todo/routes"
	"net/http"
	"os"
)



func main() {
	e := routes.Init()
	port := os.Getenv("PORT")	

	config.NewDB()

	mux := http.NewServeMux()
  mux.HandleFunc("/plm/cors",routes.Cors)

	e.Logger.Fatal(e.Start(":"+port))
}

