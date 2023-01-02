/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: app.go
  @author: Lido
  @time: 2023-01-02 16:50
  @description:
*/
package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
