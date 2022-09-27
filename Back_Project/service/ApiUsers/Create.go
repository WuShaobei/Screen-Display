package users

import (
	"Back_Project/dao"
	"Back_Project/manager"
	"Back_Project/types"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateMember(c *gin.Context) {
	var request types.CreateMemberRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	data := manager.SelectUser(request.Username)
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
	data = manager.SelectUser(request.Username)

	c.JSON(200, types.CreateMemberResponse{
		Code: types.OK,
		Data: struct{ Id uint }{Id: data.Id},
	})

}
