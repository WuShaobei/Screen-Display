package users

import (
	"Back_Project/manager"
	"Back_Project/types"
	"log"

	"github.com/gin-gonic/gin"
)

// WhoAmI
//
//	@Description: 通过 request 从数据库中获取用户数据并返回
//	@param c
//	@data 2022-09-28 14:01:32
//	@author WuShaobei
func WhoAmI(c *gin.Context) {
	var request types.WhoAmIRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	data := manager.SelectUserById(request.Id)

	c.JSON(200, types.WhoAmIResponse{
		Code: types.OK,
		Data: struct {
			Username string
			RealName string
			Identity string
		}{
			Username: data.Username,
			RealName: data.RealName,
			Identity: data.Identity,
		},
	})
}
