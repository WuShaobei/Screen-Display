package dao

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

// RDB 全局 Redis 操作变量
var RDB *redis.Client

// InitRedis
//
//	@Description: 初始化 Redis 连接
//	@data 2022-09-28 13:42:39
//	@author WuShaobei
func InitRedis() {
	// 连接 Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	// 利用根Context创建一个父Context
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := RDB.Ping().Result()

	if err != nil {
		fmt.Println("open redis fail")
		return
	}

	// 删除 redis 缓存
	res, err := RDB.FlushDB().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("delete redis:", res)

	fmt.Println("open redis success")

}
