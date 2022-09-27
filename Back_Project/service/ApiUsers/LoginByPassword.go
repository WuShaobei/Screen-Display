package users

import (
	"Back_Project/dao"
	"Back_Project/manager"
	"Back_Project/types"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

func LoginByPassword(c *gin.Context) {

	var request types.LoginRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	data := manager.SelectUser(request.Username)
	if data == (types.ChineseCateringUser{}) {
		c.JSON(200, types.LoginResponse{
			Code: types.WrongPassword,
		})
	}

	request.Password = dao.MD5(request.Password)
	if request.Password != data.Password {
		c.JSON(200, types.LoginResponse{
			Code: types.WrongPassword,
		})
	}

	// 获取唯一标识符 uuid 作为该数据行的键
	sessionKey := uuid.NewV4().String()

	if err := dao.RDB.HSet(sessionKey, "Username", data.Username).Err(); err != nil {
		log.Fatal(err)
		return
	}
	if err := dao.RDB.HSet(sessionKey, "Password", data.Password).Err(); err != nil {
		log.Fatal(err)
		return
	}
	// 设置 3 天的存储时间
	_, err := dao.RDB.Expire(sessionKey, 3*24*3600*time.Second).Result()
	if err != nil {
		return
	}

	response := types.LoginResponse{
		Code: types.OK,
		Data: struct {
			Session string
			Id      uint
		}{Session: sessionKey, Id: data.Id},
	}

	c.JSON(200, response)
}
