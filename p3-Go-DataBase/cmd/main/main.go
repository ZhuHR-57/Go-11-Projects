/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: main.go
  @author: Lido
  @time: 2023-01-02 16:49
  @description:
*/
package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"p3-Go-DataBase/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	log.Println("Starting at server 8081")

	log.Fatalln(http.ListenAndServe(":8081", r))
}
