/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: main.go
  @author: Lido
  @time: 2023-01-11 17:10
  @description:
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"forumProject/dao/mysql"
	"forumProject/dao/redis"
	"forumProject/logger"
	"forumProject/routes"
	"forumProject/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {

	// 0. flag参数指定配置文件
	var configFileName string
	flag.StringVar(&configFileName, "config", "./config.yaml", "配置文件")
	flag.Parse()

	// 1. 加载配置文件
	if err := settings.Init(configFileName); err != nil {
		fmt.Printf("init settings failed, err:%#v\n", err)
		return
	}
	fmt.Println("settings init success...")

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%#v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	/*
		3. 初始化数据库链接
	*/
	// 3.1 初始化MySQL连接（sqlx）
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	zap.L().Debug("mysql init success...")

	// 3.2 初始化Redis连接（go-redis）
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	zap.L().Debug("redis init success...")

	// 4. 路由注册
	r := routes.Setup(settings.Conf.Mode)
	zap.L().Debug("routes init success...")

	// 5. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: " + err.Error())
		}
	}()
	//
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info(fmt.Sprintf("触发关闭等待服务，将等待%ds", settings.Conf.WaitTime))
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal(fmt.Sprintf("等了%ds了还没好，先撤了...", settings.Conf.WaitTime), zap.Error(err))
	}

	zap.L().Info("所有请求处理完成,服务正常退出")
}
