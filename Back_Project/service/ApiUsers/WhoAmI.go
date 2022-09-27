package users

import (
	"Back_Project/dao"
	"Back_Project/types"
	"log"

	"github.com/gin-gonic/gin"
)

func WhoAmI(c *gin.Context) {
	var request types.WhoAmIRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	var data types.ChineseCateringUser
	dao.DB.Where("Id = ?", request.Id).Find(&data)

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
