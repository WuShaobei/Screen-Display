package users

import (
	"Back_Project/dao"
	"Back_Project/manager"
	"Back_Project/types"
	"log"

	"github.com/gin-gonic/gin"
)

// LoginBySession
//
//	@Description: 通过 request.SessionKey 从 Redis 获取 Username 和 Password 并与数据库比对，比对成功后返回 Id
//	@param c
//	@data 2022-09-28 13:57:30
//	@author WuShaobei
func LoginBySession(c *gin.Context) {

	var request types.LoginRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	username, err := dao.RDB.HGet(request.SessionKey, "Username").Result()
	if err != nil {
		c.JSON(200, types.LoginResponse{
			Code: types.UnknownError,
		})
		return
	}

	data := manager.SelectUserByUsername(username)
	if data == (types.ChineseCateringUser{}) {
		c.JSON(200, types.LoginResponse{
			Code: types.WrongPassword,
		})
		return
	}

	password, err := dao.RDB.HGet(request.SessionKey, "Password").Result()
	if err != nil {
		c.JSON(200, types.LoginResponse{
			Code: types.UnknownError,
		})
		return
	}

	if password == data.Password {
		c.JSON(200, types.LoginResponse{
			Code: types.OK,
			Data: struct {
				Session string
				Id      uint
			}{Session: "", Id: data.Id},
		})
	} else {
		c.JSON(200, types.LoginResponse{
			Code: types.WrongPassword,
		})
	}

}
