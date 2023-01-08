/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: zap.go
  @author: Lido
  @time: 2023-01-08 12:37
  @description: zap日志库的使用
*/
package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
)

var logger *zap.Logger

func main() {
	InitLogger()
	// 程序退出之前，将缓冲区中的数据刷到log文件中
	defer logger.Sync()

	// 测试日志切割
	for i := 0; i < 50000; i++ {
		logger.Info("test log split ... ")
	}

	// 记录日志
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}

//
// @Title InitLogger
// @Description 初始化zap日志库
// @Author lido 2023-01-08 14:01:48
//
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) // Debug级别

	// zap.AddCaller() 记录调用函数的信息
	logger = zap.New(core,zap.AddCaller())
}

//
// @Title getEncoder
// @Description 编码格式
// @Author lido 2023-01-08 14:01:39
// @Return zapcore.Encoder
//
func getEncoder() zapcore.Encoder {

	// 更详细的信息配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder


	// 按照json格式编码格式
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// 输出到终端的格式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//
// @Title getLogWriter
// @Description 记录位置
// @Author lido 2023-01-08 14:01:27
// @Return zapcore.WriteSyncer
//
//func getLogWriter() zapcore.WriteSyncer {
//	file, _ := os.OpenFile("../log/test.log",os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
//	return zapcore.AddSync(file)
//}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "../log/test.log",
		MaxSize:    1,  	// M
		MaxBackups: 5,   	// 最大备份数量
		MaxAge:     30, 	// 最大备份天数
		Compress:   false,  //是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
