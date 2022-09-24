package main

import (
	"Back_Project/dao"
	"Back_Project/middlewares"
	"Back_Project/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("hello World")

	//Sql
	dao.ConnectDb()
	dao.InitRedis()

	// Router
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.Use(middlewares.Cors())
	router.RegisterRouter(g)
	router.RegisterRouterTest(g)
	err := g.Run(":1432")
	if err != nil {
		return
	}
}
