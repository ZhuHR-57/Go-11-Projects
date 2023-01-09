/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: viper.go
  @author: Lido
  @time: 2023-01-08 15:19
  @description: viper使用
*/
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/spf13/pflag"

	"github.com/gin-gonic/gin"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Port        string `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	DbName string `mapstructure:"dbname"`
	User   string `mapstructure:"user"`
	Pwd    string `mapstructure:"pwd"`
}

func initSqlXDB() (err error) {

	var c Config

	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
	}

	// dsn := "root:rootroot@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		c.MySQLConfig.User,
		c.MySQLConfig.Pwd,
		c.MySQLConfig.Host,
		c.MySQLConfig.Port,
		c.MySQLConfig.DbName,
	)
	fmt.Println(dsn)

	// 也可以使用MustConnect连接不成功就panic
	sql_db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	sql_db.SetMaxOpenConns(20)
	sql_db.SetMaxIdleConns(10)
	return
}

func main() {

	// 使用标准库 "flag" 包
	// flag.Int("flagname", 1234, "help message for flagname")
	flag.Int("age", 1138, "Port to run Application server on")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	fmt.Println(viper.GetInt("age")) // 从 viper 检索值

	// 设置默认值 (优先级最底)
	viper.SetDefault("fireDir", "./")

	// 直接写明配置文件
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径

	// 详细的配置文件
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".") // 还可以在工作目录中查找配置

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Not Found config! ")
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	if err := initSqlXDB(); err != nil {
		panic(err)
	}
	fmt.Println("Connect to Mysql Succeed!")

	r := gin.Default()
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})
	r.Run(viper.GetString("port"))
}
