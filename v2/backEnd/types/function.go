package types

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func Decimal(value float64) string {
	return fmt.Sprintf("%.2f", value+0.005)
}

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

func GetIdentity(identity IdentityTypes) IdentityTypes {
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

func GetTime(year, month int) string {
	return strconv.Itoa(year) + "-" + strconv.Itoa(month)
}
