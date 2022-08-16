/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: main.go
  @author: Lido
  @time: 2022-08-16 11:51
  @description: 主程序入口
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

//
// @Title formHandle
// @Description chu
// @Author lido 2022-08-16 12:02:27
// @Param writer
// @Param request 来自用户的请求URL
//
func formHandle(writer http.ResponseWriter, request *http.Request) {

	//处理表单
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err :%v", err)
		return
	}

	fmt.Fprintf(writer, "POST request successful!\n")

	//获取表单参数
	name := request.FormValue("name")
	address := request.FormValue("address")

	fmt.Fprintf(writer, "NAME is %v", name)
	fmt.Fprintf(writer, "Addr is %v", address)
}

//
// @Title helloHandle
// @Description
// @Author lido 2022-08-16 12:22:22
// @Param writer
// @Param request
//
func helloHandle(writer http.ResponseWriter, request *http.Request) {

	//判断用户的URL
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}

	//判断用户的请求方法
	if request.Method != "GET" {
		http.Error(writer, "method is not support", http.StatusNotFound)
		return
	}

	fmt.Fprintf(writer, "hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./p1-Go-Server/static"))
	http.Handle("/", fileServer)

	//处理from路由
	http.HandleFunc("/form", formHandle)
	//处理hello路由
	http.HandleFunc("/hello", helloHandle)

	fmt.Printf("Start Sever at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
