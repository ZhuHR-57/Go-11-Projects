/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: redisDemo.go
  @author: Lido
  @time: 2023-01-07 13:10
  @description: 连接redis
*/
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var rdb *redis.Client
var ctx,cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
//var ctx = context.Background()

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "47.107.52.134:6379",
		Password: "root", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})

	_, err = rdb.Ping(ctx).Result()

	return err
}

func GetAndSet(){

	// set
	err := rdb.Set(ctx, "name", "lido", 0).Err()
	if err != nil {
		panic(err)
	}

	//get
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key:name is", val)

	// 当获取不存在的key时
	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		fmt.Printf("Failed to Get err:%#v",err)
	} else {
		fmt.Println("key2", val2)
	}
}

// zsetDemo 操作zset示例
func zsetDemo() {
	// key
	zsetKey := "language_rank"

	// value
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}

	// ZADD,添加数据
	err := rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 更新单个值
	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

// watchDemo 在key值不变的情况下将其值+1
func watchDemo(ctx context.Context, key string) error {
	return rdb.Watch(ctx, func(tx *redis.Tx) error {
		_,err := tx.Get(ctx, key).Result()
		if err != nil && err != redis.Nil {
			return err
		}
		// 假设操作耗时5秒
		// 5秒内我们通过其他的客户端修改key，当前事务就会失败
		time.Sleep(60 * time.Second)
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			//pipe.Set(ctx, key, n+1, time.Hour)
			fmt.Println("居然没被修改...")
			return nil
		})
		return err
	}, key)
}

func main() {

	if err := initClient(); err != nil {
		log.Fatalln("init redis failed", err)
	}

	fmt.Println("Connect Succeed!")
	defer rdb.Close()
	defer cancel()

	if err := watchDemo(ctx,"name");err != nil{
		fmt.Println(err)
	}

}
