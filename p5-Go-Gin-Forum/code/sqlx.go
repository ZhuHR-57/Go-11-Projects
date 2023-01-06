/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: sqlx.go
  @author: Lido
  @time: 2023-01-06 12:42
  @description:
*/
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var sql_db *sqlx.DB

type USER struct {
	Id   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

func initSqlXDB() (err error) {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	sql_db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	sql_db.SetMaxOpenConns(20)
	sql_db.SetMaxIdleConns(10)
	return
}

func queryMultiRowSqlx() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []USER
	err := sql_db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	for _,v := range users{
		fmt.Println(v.Name)
	}

	fmt.Printf("users:%#v\n", users)
}

func insertUser()(err error){
	// :name,:age
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	_, err = sql_db.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "XiaoLi",
			"age": 28,
		})
	return
}

func main() {
	if err := initSqlXDB(); err != nil {
		panic(err)
	}
	fmt.Println("Connect to Mysql Succeed!")

	insertUser()
}
