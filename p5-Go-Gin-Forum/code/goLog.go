/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: goLog.go
  @author: Lido
  @time: 2023-01-08 12:41
  @description: 使用Go原生的日志库
*/
package main

import (
	"log"
	"net/http"
	"os"
)

func SetupLogger() {
	logFileLocation, _ := os.OpenFile("../log/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}

func simpleHttpGet_golog(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
}

func main() {
	SetupLogger()
	simpleHttpGet_golog("www.baidu.com")
	simpleHttpGet_golog("http://www.baidu.com")
}
