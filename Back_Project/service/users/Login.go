package users

import (
	"Back_Project/dao"
	"Back_Project/types"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func Login(c *gin.Context) {
	// 获取请求
	var request types.LoginRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	// 获取数据行
	var line types.ChineseCateringUser
	dao.DB.Model(&types.ChineseCateringUser{}).Where(&request).Find(&line)
	if line == (types.ChineseCateringUser{}) {
		c.JSON(200, types.LoginResponse{
			Code: types.WrongPassword,
		})
		return
	}

	// 获取唯一标识符 uuid 作为该数据行的键
	sessionKey := uuid.NewV4().String()

	// redis 记录 sessionKey 对应的 UserID， UserType

	// TODO 优化 UserType 写法
	data := map[string]interface{}{
		"UserID":   line.Id,
		"UserType": fmt.Sprint(line.Identity),
	}
	if err := dao.RDB.HMSet(sessionKey, data).Err(); err != nil {
		log.Fatal(err)
		return
	}

	// 设置 cookie
	c.SetCookie("camp-session", sessionKey, 3600, "/", "", false, true)

	response := types.LoginResponse{
		Code: types.OK,
		Data: struct{ Id string }{
			Id: line.Id,
		},
	}

	c.JSON(200, response)
}
