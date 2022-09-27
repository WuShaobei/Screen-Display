package users

import (
	"Back_Project/dao"
	"Back_Project/manager"
	"Back_Project/types"
	"log"

	"github.com/gin-gonic/gin"
)

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

	data := manager.SelectUser(username)
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
