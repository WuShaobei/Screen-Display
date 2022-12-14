// Package service
// @Description: 薪资相关接口

package service

import (
	"backEnd/manager"
	"backEnd/types"
	"log"

	"github.com/gin-gonic/gin"
)

// PostAvgSalaryAndCountsByYearFromPaymentService
//
//	@Description: 获取平均薪资和用户类型接口
//	@param c
func PostAvgSalaryAndCountsByYearFromPaymentService(c *gin.Context) {
	var request types.PostAmountAndPercentageByYearFromStatisticRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Println(err)
		return
	}

	if request.BeginYear < types.PaymentMinYear {
		request.BeginYear = types.PaymentMinYear
	}

	if request.EndYear < types.PaymentMaxYear {
		request.EndYear = types.PaymentMaxYear
	}

	paymentManager := manager.PaymentManager{}

	if data, errNo := paymentManager.PostAvgSalaryAndCountsByYearFromPaymentManage(request.BeginYear, request.EndYear); errNo != types.OK {
		c.JSON(200, types.PostAvgSalaryAndCountsByYearFromPaymentResponse{Code: errNo})
	} else {
		c.JSON(200, types.PostAvgSalaryAndCountsByYearFromPaymentResponse{Code: errNo, Data: data})
	}
}
