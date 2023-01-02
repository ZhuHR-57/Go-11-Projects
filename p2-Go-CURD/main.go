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
	"math/rand"
	"net/http"
	"strconv"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var books []Book

//
// @Title initBooks
// @Description 创建一些书籍，用于测试
// @Author lido 2023-01-02 11:18:10
//
func initBooks() {
	books = append(books, Book{
		ID:    "1",
		Isbn:  "1234",
		Title: "Book Title 1",
		Author: &Author{
			Name: "author 1",
			Age:  1,
		},
	})
	books = append(books, Book{
		ID:    "2",
		Isbn:  "1235",
		Title: "Book Title 2",
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

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getbook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Print("Starting server at port 8000")

	log.Fatal(http.ListenAndServe(":8000", r))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
	return
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	//1. 设置json头
	w.Header().Set("Content-Type", "application/json")
	//2. 获取参数
	params := mux.Vars(r)
	//3. 找到需要删除的ID并删除
	for index, item := range books {
		if item.ID == params["id"] {
			// 删除指定的元素
			books = append(books[:index], books[index+1:]...)
			return
		}
	}
	return
}

func getbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	return
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

	return
}

func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// 找到元素并删除
	for index, item := range books {
		if item.ID == params["id"] {
			// 1. 删除原来的书
			books = append(books[:index], books[index+1:]...)
			// 2. 建立新的书
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] // 原来的ID不变
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	return
}
