package manager

import (
	"Back_Project/dao"
	"Back_Project/types"
)

func SelectUser(Username string) types.ChineseCateringUser {
	var data types.ChineseCateringUser

	dao.DB.Where("username=?", Username).Find(&data)

	return data
}

func InsertUser(data types.ChineseCateringUser) {
	dao.DB.Create(&data)
}
