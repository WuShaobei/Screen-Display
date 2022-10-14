// Package types
// @Description: 辅助函数
package types

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

// Decimal
//
//	@Description: 四舍五入保留两位小数并转化成字符串
//	@param value float64
//	@return string
func Decimal(value float64) string {
	return fmt.Sprintf("%.2f", value+0.005)
}

// MD5
//
//	@Description: 字符串转化成对应的 MD5码字符串以实现 加密功能
//	@param str
//	@return string
func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

// GetIdentity
//
//	@Description: 用户类型代码转换为其全称
//	@param identity
//	@return string
func GetIdentity(identity IdentityTypes) string {
	if identity == Admin {
		return "管理员"
	}
	if identity == Investor {
		return "投资人"
	}
	if identity == Practitioner {
		return "从业者"
	}
	if identity == Tourist {
		return "游客"
	}
	return ""
}

// GetTime
//
//	@Description: 将 year 和 month 转化为 year-month 格式的字符串
//	@param year int
//	@param month int
//	@return string
func GetTime(year, month int) string {
	return strconv.Itoa(year) + "-" + strconv.Itoa(month)
}
