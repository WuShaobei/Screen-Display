package users

import (
	"Back_Project/dao"
	"Back_Project/manager"
	"Back_Project/types"
	"log"

	"github.com/gin-gonic/gin"
)

// Register
//
//	@Description: 从 request 中获取用户数据，然后插入数据库中，提取用户 Id 并返回
//	@param c
//	@data 2022-09-28 14:00:42
//	@author WuShaobei
func Register(c *gin.Context) {
	var request types.CreateMemberRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	data := manager.SelectUserByUsername(request.Username)
	if data != (types.ChineseCateringUser{}) {
		c.JSON(200, types.CreateMemberResponse{Code: types.UserHasExisted})
		return
	}

	manager.InsertUser(types.ChineseCateringUser{
		Username: request.Username,
		Password: dao.MD5(request.Password),
		RealName: request.RealName,
		Identity: request.Identity,
	})
	data = manager.SelectUserByUsername(request.Username)

	c.JSON(200, types.CreateMemberResponse{
		Code: types.OK,
		Data: struct{ Id uint }{Id: data.Id},
	})

}
