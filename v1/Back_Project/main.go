package main

import (
	"Back_Project/dao"
	"Back_Project/middlewares"
	"Back_Project/router"

	"github.com/gin-gonic/gin"
)

func main() {

	//Sql
	dao.ConnectDb()
	// TODO 判断是否要重置数据表
	dao.InitTables()

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
