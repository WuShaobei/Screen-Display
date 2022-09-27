package dao

import (
	"Back_Project/types"
	"crypto/md5"
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	userName = "root"
	passWord = "12345678"
	ip       = "localhost"
	port     = "3306"
	dbName   = "Final_Project"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := strings.Join([]string{userName, ":", passWord, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8mb4&parseTime=True"}, "")

	fmt.Println("Connecting " + dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		fmt.Println("Open Database Fail")
		return
	}
	sqlDb, _ := DB.DB()
	// 设置空闲连接数
	// 数量 connections = ((core_count * 2) + effective_spindle_count)
	sqlDb.SetConnMaxIdleTime(10)
	// 最大连接数
	sqlDb.SetMaxOpenConns(100)
	// 连接复用连接时间
	sqlDb.SetConnMaxLifetime(time.Hour)
	fmt.Println("Mysql Connect Success")
}

/**
 * 创建销售表
 *
 */

// MD5
// 密码加密
func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

func InitUserTable() {
	// 删除 user
	if err := DB.Exec("DROP TABLE chinese_catering_users"); err != nil {
		fmt.Println(err)
	}
	if err := DB.AutoMigrate(&types.ChineseCateringUser{}); err != nil {
		return
	}
	DB.Create(&types.ChineseCateringUser{
		Username: "LocalHost",
		Password: MD5("127.0.0.1"),
		RealName: "Admin",
		Identity: types.Admin,
	})
	DB.Create(&types.ChineseCateringUser{
		Username: "Investor",
		Password: MD5("Investor"),
		RealName: "Investor",
		Identity: types.Investor,
	})
	DB.Create(&types.ChineseCateringUser{
		Username: "Practitioner",
		Password: MD5("Practitioner"),
		RealName: "Practitioner",
		Identity: types.Practitioner,
	})
	DB.Create(&types.ChineseCateringUser{
		Username: "Tourist",
		Password: MD5("Tourist"),
		RealName: "Tourist",
		Identity: types.Tourist,
	})
	fmt.Println("create table success")
}
