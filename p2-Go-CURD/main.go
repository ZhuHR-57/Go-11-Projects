/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: main.go
  @author: Lido
  @time: 2023-01-02 10:12
  @description:
*/
package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

var books []Book

//
// @Title initBooks
// @Description 创建一些书籍，用于测试
// @Author lido 2023-01-02 11:18:10
//
func initBooks(){
	books = append(books,Book{
		ID:     "1",
		Isbn:   "1234",
		Title:  "Book Title 1",
		Author: &Author{
			Name: "author 1",
			Age:  1,
		},
	})
	books = append(books,Book{
		ID:     "2",
		Isbn:   "1235",
		Title:  "Book Title 2",
		Author: &Author{
			Name: "author 2",
			Age:  2,
		},
	})

}

func main() {

	//1. 初始化一些测试数据
	initBooks()

	//多路复用
	r := mux.NewRouter()

	r.HandleFunc("/books",getBooks).Methods("GET")
	r.HandleFunc("/books/{id}",getbook).Methods("GET")
	r.HandleFunc("/books",createBook).Methods("POST")
	r.HandleFunc("/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}",deleteBook).Methods("DELETE")

	log.Print("Starting server at port 8080")

	if err := http.ListenAndServe(":8081",nil); err != nil{
		log.Fatal(err)
	}

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)


}

