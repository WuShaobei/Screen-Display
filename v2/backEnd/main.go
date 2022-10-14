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

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	/// 初始化用户表 删除所有用户 然后新增管理员用户
	/// 	* username : administration
	/// 	* password : administration
	/// 	* realName : admin
	/// 	* identity : 管理员
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	///if err := dao.DB.Exec("DROP TABLE chinese_catering_users"); err != nil {//
	///	fmt.Println(err)
	///}
	///if err := dao.DB.AutoMigrate(&types.ChineseCateringUser{}); err != nil {
	///	return
	///}
	///
	///userDao := dao.UserDao{}
	///if _, errNo := userDao.InsertUserDataToMySQL("administration", types.MD5("administration"), "admin", types.Admin); errNo == types.OK {
	///	fmt.Println("InitUserTable Success")
	///} else {
	///	fmt.Println("InitUserTable Err")
	///}
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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
