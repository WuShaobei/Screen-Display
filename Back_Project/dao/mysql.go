package dao

import (
	"fmt"
	"server_side/types"
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
	dbName   = "project_9_20"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := strings.Join([]string{userName, ":", passWord, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8mb4&parseTime=True"}, "")

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		fmt.Println("open database fail")
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
	fmt.Println("connect success")
}

/**
 * 创建销售表
 *
 */

func InitSalesVolumeTable() {
	if err := DB.Exec("DROP TABLE sales_volumes"); err != nil {
		fmt.Println(err)
	}
	if err := DB.AutoMigrate(&types.SalesVolume{}); err != nil {
		return
	}
	fmt.Println("create table success")
}
