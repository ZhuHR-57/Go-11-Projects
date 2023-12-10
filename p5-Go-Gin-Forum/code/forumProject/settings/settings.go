/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: settings.go
  @author: Lido
  @time: 2023-01-11 17:28
  @description: 配置文件的读取
*/
package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    uint16 `mapstructure:"machine_id"`
	WaitTime     int    `mapstructure:"wait_time"`
	Salt         string `mapstructure:"salt"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init(configFileName string) (err error) {

	// 1.相对路径（是相对于执行的位置）
	viper.SetConfigFile(configFileName)

	// 2.注意不要重名
	//viper.SetConfigName("config")
	//viper.AddConfigPath(".")
	//viper.AddConfigPath("./conf")

	// 3.远程配置中心获取 使用什么格式去解析
	//viper.SetConfigType("yaml") // json

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Not Found config! ")
		} else {
			fmt.Printf("Fatal error config file: %s \n", err)
		}
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	viper.WatchConfig()
	// 小回调的钩子
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了，重新加载到全局Conf ...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})

	return nil
}
