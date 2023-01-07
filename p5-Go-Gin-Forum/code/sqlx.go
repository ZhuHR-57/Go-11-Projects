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
	"database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

var sql_db *sqlx.DB

type USER struct {
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

func (u USER) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

func BatchInsertUsers(users []*USER) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users) * 2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := sql_db.Exec(stmt, valueArgs...)
	return err
}

// BatchInsertUsers_Sqlx_in 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers_Sqlx_in(users []interface{}) error {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?), (?)", // 插入三条数据
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args
	_, err := sql_db.Exec(query, args...)
	return err
}

// BatchInsertUsers3_NameExec 使用NamedExec实现批量插入
func BatchInsertUsers3_NameExec(users []*USER) error {
	_, err := sql_db.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}

func queryMultiRow() {
	sqlStr := "select name, age from user where id > ?"
	var users []USER
	err := sql_db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	for _,v := range users{
		fmt.Println(v.Name) // 需要的字段
	}

	fmt.Printf("users:%#v\n", users)
}

func main() {
	if err := initSqlXDB(); err != nil {
		panic(err)
	}
	fmt.Println("Connect to Mysql Succeed!")

	u1 := USER{Name: "test1", Age: 18}
	u2 := USER{Name: "test2", Age: 28}
	u3 := USER{Name: "test3", Age: 38}

	users := []*USER{&u1, &u2, &u3}
	err := BatchInsertUsers(users)
	if err != nil {
		fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
	}

	u4 := USER{Name: "test4", Age: 18}
	u5 := USER{Name: "test5", Age: 28}
	u6 := USER{Name: "test6", Age: 38}

	// 方法2
	users2 := []interface{}{u4, u5, u6}
	err = BatchInsertUsers_Sqlx_in(users2)
	if err != nil {
		fmt.Printf("BatchInsertUsers2 failed, err:%v\n", err)
	}

	u7 := USER{Name: "test7", Age: 18}
	u8 := USER{Name: "test8", Age: 28}
	u9 := USER{Name: "test9", Age: 38}

	// 方法3
	users3 := []*USER{&u7, &u8, &u9}
	err = BatchInsertUsers3_NameExec(users3)
	if err != nil {
		fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)
	}


	queryMultiRow()
}
