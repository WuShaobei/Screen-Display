package manager

import (
	"Back_Project/dao"
	"Back_Project/types"
)

// SelectUserByUsername
//
//	@Description: 根据用户名查找用户
//	@param Username
//	@return types.ChineseCateringUser
//	@data 2022-09-28 13:53:32
//	@author WuShaobei
func SelectUserByUsername(Username string) types.ChineseCateringUser {
	var data types.ChineseCateringUser

	dao.DB.Where("username=?", Username).Find(&data)

	return data
}

// SelectUserById
//
//	@Description: 根据用户 Id 查找用户
//	@param Id
//	@return types.ChineseCateringUser
//	@data 2022-09-28 13:53:51
//	@author WuShaobei
func SelectUserById(Id string) types.ChineseCateringUser {
	var data types.ChineseCateringUser

	dao.DB.Where("Id=?", Id).Find(&data)

	return data
}

// InsertUser
//
//	@Description: 插入用户数据
//	@param data
//	@data 2022-09-28 13:54:43
//	@author WuShaobei
func InsertUser(data types.ChineseCateringUser) {
	dao.DB.Create(&data)
}
