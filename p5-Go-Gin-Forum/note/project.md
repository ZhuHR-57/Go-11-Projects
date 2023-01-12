# 论坛架构

![image-20230111170521897](project.assets/image-20230111170521897.png)

# 项目框架搭建

> 0. IDE设置

![image-20230111171640532](project.assets/image-20230111171640532.png)

1. 配置代理 和 Go Modules

```
GOPROXY=https://goproxy.cn,direct
```

> 2. 项目框架

```go
func main() {
	
	// 1. 加载配置文件
	
	// 2. 初始化日志
	
	// 3. 初始化数据库链接
	
	// 4. 路由注册
	
	// 5. 启动服务（优雅关机）

}
```

> 3. 加载配置文件

`配置文件`

```yaml
name: "web_app"
mode: "dev"
port: 8080
version: "v0.1.4"

# 雪花算法：开始时间 机器ID
start_time: "2020-07-01"
machine_id: 1

log:
  level: "debug"
  filename: "web_app.log"
  max_size: 200
  max_age: 30
  max_backups: 7
mysql:
  host: "127.0.0.1"
  port: 3306
  user: "root"
  password: "chnsys@2016"
  dbname: "gin_bbs"
  max_open_conns: 200
  max_idle_conns: 50
redis:
  host: "127.0.0.1"
  port: 6379
  password: ""
  db: 0
  pool_size: 100
```

`加载配置文件`

```

```

