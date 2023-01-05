/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: main.go
  @author: Lido
  @time: 2023-01-05 9:53
  @description:
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func func1(c *gin.Context){
	fmt.Println("func1")
	c.Set("name","lido")
}

func func2(c *gin.Context){
	fmt.Println("func2 before")
	c.Next()
	fmt.Println("func2 after")
}

func func3(c *gin.Context){
	fmt.Println("func3")
	v,ok := c.Get("name")
	if ok {
		name := v.(string)
		fmt.Println(name)
	}
	c.Abort()
}

func func4(c *gin.Context){
	fmt.Println("func4")
}

func func5(c *gin.Context){
	fmt.Println("func5")
}

func main() {

	r := gin.Default()

	oneGroup := r.Group("/hello", func1,func2)
	oneGroup.Use(func3)

	{
		oneGroup.GET("/index",func4,func5)
	}

	r.Run()
}
