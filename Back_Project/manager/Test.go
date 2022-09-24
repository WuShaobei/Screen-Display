package manager

import (
	"Back_Project/dao"
	"Back_Project/types"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type MyObject struct{}

type TestRequest struct {
	Table int
}

func TestApiSelectFirst(c *gin.Context) {
	var request TestRequest

	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	fmt.Println(request)
	if request.Table == types.Table1ChineseCateringBrandStatistic {
		data := types.ChineseCateringBrandStatistic{}
		dao.DB.First(&data)
		log.Println(data)
		c.JSON(200, data)
		return
	}
	if request.Table == types.Table2ChineseCateringFundingStatistic {
		data := types.ChineseCateringFundingStatistic{}
		dao.DB.First(&data)

		log.Println(data)
		c.JSON(200, data)
		return
	}
	if request.Table == types.Table3ChineseCateringOnlineOrderStatistic {
		data := types.ChineseCateringOnlineOrderStatistic{}
		dao.DB.First(&data)

		log.Println(data)
		c.JSON(200, data)
		return
	}
	if request.Table == types.Table4ChineseCateringPayment {
		data := types.ChineseCateringPayment{}
		dao.DB.First(&data)

		log.Println(data)
		c.JSON(200, data)
		return
	}
	if request.Table == types.Table5ChineseCateringStatistic {
		data := types.ChineseCateringStatistic{}
		dao.DB.First(&data)

		log.Println(data)
		c.JSON(200, data)
		return
	}
	if request.Table == types.Table6ChineseCateringUser {
		data := types.ChineseCateringUser{}
		dao.DB.First(&data)

		log.Println(data)
		c.JSON(200, data)
		return
	}

}
