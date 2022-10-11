package types

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value+0.005), 64)
	return value
}

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
