/*
 * @Author: WuShaobei
 * @Date: 2022-10-09 21:00:27
 * @LastEditTime: 2022-10-09 21:01:44
 * @FilePath: /Screen-Display/v2/backEnd/main.go
 * @Description: 主函数
 */
package main

import (
	"backEnd/dao"
	"backEnd/middlewares"
	"backEnd/router"

	"github.com/gin-gonic/gin"
)

func main() {

	// 连接数据库
	dao.ConnectDb()

	/// 重置 User 数据表
	///dao.InitUserTable()

	// 连接 Redis
	dao.InitRedis()

	// Router
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.Use(middlewares.Cors())
	router.RegisterRouter(g)

	err := g.Run(":1432")
	if err != nil {
		return
	}
}
