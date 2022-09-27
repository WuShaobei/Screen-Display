package users

import (
	"Back_Project/dao"
	"Back_Project/types"
	"fmt"
	"log"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// 获取请求 Username, Password | SessionKey
	var request types.LoginRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}
	fmt.Println(request)
	// 通过 sessionKey 登录
	if request.SessionKey != "" {
		// 获取用户名和密码
		username, err := dao.RDB.HGet(request.SessionKey, "Username").Result()
		if err != nil {

			c.JSON(200, types.LoginResponse{
				Code: types.UnknownError,
			})
			panic(err)
			return
		}
		password, err := dao.RDB.HGet(request.SessionKey, "Password").Result()
		if err != nil {

			c.JSON(200, types.LoginResponse{
				Code: types.UnknownError,
			})
			panic(err)
			return
		}

		// 获取用户数据
		var line types.ChineseCateringUser
		dao.DB.Model(&types.ChineseCateringUser{}).Where("username = ? && password = ?", username, password).Find(&line)
		if line == (types.ChineseCateringUser{}) {
			c.JSON(200, types.LoginResponse{
				Code: types.WrongPassword,
			})
			return
		}

		// 返回参数，结束函数
		c.JSON(200, types.LoginResponse{
			Code: types.OK,
			Data: line,
		})
		return
	}

	request.Password = dao.MD5(request.Password)
	// 获取数据行
	var line types.ChineseCateringUser
	dao.DB.Model(&types.ChineseCateringUser{}).Where("username = ? && password = ?", request.Username, request.Password).Find(&line)
	if line == (types.ChineseCateringUser{}) {
		c.JSON(200, types.LoginResponse{
			Code: types.WrongPassword,
		})
		return
	}

	// 获取唯一标识符 uuid 作为该数据行的键
	sessionKey := uuid.NewV4().String()

	// redis 记录 sessionKey 对应的 Id，Username Password, Identity
	if err := dao.RDB.HSet(sessionKey, "Id", line.Id).Err(); err != nil {
		log.Fatal(err)
		return
	}
	if err := dao.RDB.HSet(sessionKey, "Username", line.Username).Err(); err != nil {
		log.Fatal(err)
		return
	}
	if err := dao.RDB.HSet(sessionKey, "Password", line.Password).Err(); err != nil {
		log.Fatal(err)
		return
	}
	if err := dao.RDB.HSet(sessionKey, "Identity", strconv.Itoa(int(line.Identity))).Err(); err != nil {
		log.Fatal(err)
		return
	}
	// 设置 3 天的存储时间
	_, err := dao.RDB.Expire(sessionKey, 3*24*3600*time.Second).Result()
	if err != nil {
		return
	}

	response := types.LoginResponse{
		Code:    types.OK,
		Session: sessionKey,
		Data:    line,
	}

	c.JSON(200, response)
}
