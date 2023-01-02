/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: bookstore-routes.go
  @author: Lido
  @time: 2023-01-02 16:52
  @description:
*/
package routes

import (
	"github.com/gorilla/mux"
	"p3-Go-DataBase/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
