package dao

import (
	"Back_Project/types"
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
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
	dbName   = "test"
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

func InitTables() {
	// 删除 数据表
	if err := DB.Exec("DROP TABLE chinese_catering_statistics"); err != nil {
		fmt.Println(err)
	}
	if err := DB.Exec("DROP TABLE chinese_catering_online_order_statistics"); err != nil {
		fmt.Println(err)
	}
	if err := DB.Exec("DROP TABLE chinese_catering_payments"); err != nil {
		fmt.Println(err)
	}
	if err := DB.Exec("DROP TABLE chinese_catering_brand_statistics"); err != nil {
		fmt.Println(err)
	}
	if err := DB.Exec("DROP TABLE chinese_catering_funding_statistics"); err != nil {
		fmt.Println(err)
	}
	if err := DB.Exec("DROP TABLE chinese_catering_users"); err != nil {
		fmt.Println(err)
	}

	// 重新创建数据表
	if err := DB.AutoMigrate(&types.ChineseCateringStatistic{}); err != nil {
		return
	}
	if err := DB.AutoMigrate(&types.ChineseCateringOnlineOrderStatistic{}); err != nil {
		return
	}
	if err := DB.AutoMigrate(types.ChineseCateringPayment{}); err != nil {
		return
	}
	if err := DB.AutoMigrate(types.ChineseCateringBrandStatistic{}); err != nil {
		return
	}
	if err := DB.AutoMigrate(&types.ChineseCateringFundingStatistic{}); err != nil {
		return
	}
	if err := DB.AutoMigrate(&types.ChineseCateringUser{}); err != nil {
		return
	}

	initChineseCateringStatistic()
	initChineseCateringOnlineOrderStatistic()
	initChineseCateringPayment()
	initChineseCateringBrandStatistic()
	initChineseCateringFundingStatistic()
	initUSer()

	fmt.Println("create table success")
}

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

func initChineseCateringStatistic() {
	for year := 2014; year <= 2021; year++ {
		DB.Create(&types.ChineseCateringStatistic{
			Year:                  year,
			TotalAmount:           Decimal(float64(rand.Intn(500)) / 10.0),
			TotalAmountPercentage: Decimal(float64(rand.Intn(200))/10.0 - 10),
		})
	}
}

func initChineseCateringOnlineOrderStatistic() {
	for year := 2010; year <= 2021; year++ {
		for month := 1; month <= 12; month++ {
			DB.Create(&types.ChineseCateringOnlineOrderStatistic{
				Year:        year,
				Month:       month,
				OrderAmount: rand.Intn(50),
			})
		}
	}
}
func initChineseCateringPayment() {
	user := 0
	for year := 2018; year <= 2021; year++ {
		for num := 0; num < 1000+rand.Intn(1000); num++ {
			user += 1
			DB.Create(&types.ChineseCateringPayment{
				Year:   year,
				Name:   "User" + strconv.Itoa(user),
				Salary: 2000 + rand.Intn(8000),
			})
		}
	}
}

func initChineseCateringBrandStatistic() {
	brand := 0
	for num := 0; num < 5+rand.Intn(5); num++ {
		brand += 1
		DB.Create(&types.ChineseCateringBrandStatistic{
			Brand: "brand" + strconv.Itoa(brand),
			Price: rand.Intn(79),
		})
	}
	for num := 0; num < 5+rand.Intn(5); num++ {
		brand += 1
		DB.Create(&types.ChineseCateringBrandStatistic{
			Brand: "brand" + strconv.Itoa(brand),
			Price: 81 + rand.Intn(39),
		})
	}
	for num := 0; num < 5+rand.Intn(5); num++ {
		brand += 1
		DB.Create(&types.ChineseCateringBrandStatistic{
			Brand: "brand" + strconv.Itoa(brand),
			Price: 121 + rand.Intn(79),
		})
	}
	for num := 0; num < 5+rand.Intn(5); num++ {
		brand += 1
		DB.Create(&types.ChineseCateringBrandStatistic{
			Brand: "brand" + strconv.Itoa(brand),
			Price: 201 + rand.Intn(800),
		})
	}
}
func initChineseCateringFundingStatistic() {
	for year := 2021; year <= 2022; year++ {
		for month := 1; month <= 12; month++ {
			for i := 0; i < rand.Intn(10); i++ {
				DB.Create(&types.ChineseCateringFundingStatistic{
					FundingYear:  year,
					FundingMonth: month,
					FundingRound: strconv.Itoa(rand.Intn(10)),
					Brand:        "Brand" + strconv.Itoa(rand.Intn(100)),
					Investor:     "Investor" + strconv.Itoa(rand.Intn(10)),
				})
			}
		}
	}
}
func initUSer() {
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
}
