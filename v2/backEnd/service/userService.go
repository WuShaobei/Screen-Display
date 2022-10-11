package service

import "C"
import (
	"backEnd/manager"

	"backEnd/types"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func WhoAmIService(c *gin.Context) {
	var request types.WhoAmIRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	var userManage manager.UserManager
	data, errNo := userManage.WhoAmI(request.Id)
	fmt.Println(data)
	if data != (types.WhoAmIData{}) {
		c.JSON(200, types.WhoAmIResponse{Code: errNo, Data: data})
	} else {
		c.JSON(200, types.WhoAmIResponse{Code: types.UserNotExist})
	}
}

func LoginByPasswordService(c *gin.Context) {
	var request types.LoginByPasswordRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	request.Password = types.MD5(request.Password)
	userManager := manager.UserManager{}
	data, errNo := userManager.LoginByPassword(request.Username, request.Password, request.SetCookie)

	if errNo == types.OK {
		c.JSON(200, types.LoginResponse{Code: types.OK, Data: data})
	} else {
		c.JSON(200, types.LoginResponse{Code: errNo})
	}

}

func LoginBySessionKeyService(c *gin.Context) {
	var request types.LoginBySessionKeyRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	userManager := manager.UserManager{}
	data, errNo := userManager.LoginBySessionKey(request.SessionKey)
	if errNo == types.OK {
		c.JSON(200, types.LoginResponse{
			Code: errNo,
			Data: data,
		})
	} else {
		c.JSON(200, types.LoginResponse{
			Code: errNo,
		})
	}
}

func LogoutService(c *gin.Context) {
	var request types.LogoutRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	userManager := manager.UserManager{}
	errNo := userManager.Logout(request.SessionKey)

	c.JSON(200, types.LogoutResponse{Code: errNo})
}

func RegisterService(c *gin.Context) {
	var request types.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	request.Password = types.MD5(request.Password)
	userManager := manager.UserManager{}
	data, errNo := userManager.Register(request)

	if errNo == types.OK {
		c.JSON(200,
			types.RegisterResponse{
				Code: types.OK,
				Data: struct {
					Id string `json:"id"`
				}{Id: data},
			},
		)
	} else {
		c.JSON(200, types.RegisterResponse{Code: types.OK})
	}
}
