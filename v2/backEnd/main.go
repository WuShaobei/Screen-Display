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

	//Sql
	dao.ConnectDb()
	// TODO 判断是否要重置数据表
	//dao.InitTables()

	// Redis
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
