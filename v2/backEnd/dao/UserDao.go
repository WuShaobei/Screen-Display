// Package dao
// @Description: ChineseCateringUser 类 与 mysql 和 redis 交互
package dao

import (
	"backEnd/types"
	"log"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

// UserDao
// @Description: 交互类
type UserDao struct {
	//user types.ChineseCateringUser
}

// SelectDataByIdFromMySQL
//
//	@Description: 通过 Id 从 mysql 获取用户信息
//	@receiver u
//	@param id
//	@return types.WhoAmIData
//	@return types.ErrNo
//	@data 2022-10-10 09:34:00
//	@author WuShaobei
func (u *UserDao) SelectDataByIdFromMySQL(id string) (types.WhoAmIData, types.ErrNo) {

	var user types.ChineseCateringUser
	DB.Where("id = ?", id).Find(&user)

	if (user != types.ChineseCateringUser{}) {
		return types.WhoAmIData{
			Id: strconv.Itoa(int(user.Id)), Username: user.Username, RealName: user.RealName, Identity: user.Identity,
		}, types.OK
	} else {
		return types.WhoAmIData{}, types.UserNotExist
	}
}

// SelectIdAndPasswordByUsernameFromMySQL
//
//	@Description: 通过 username 从 Mysql 读取 Id 和 password
//	@receiver u
//	@param username
//	@data 2022-10-10 09:34:54
//	@author WuShaobei
func (u *UserDao) SelectIdAndPasswordByUsernameFromMySQL(username string) (
	struct {
		Id       string
		Password string
	}, types.ErrNo,
) {
	user := types.ChineseCateringUser{}

	DB.Where("username = ?", username).Find(&user)

	if (user == types.ChineseCateringUser{}) {
		return struct {
			Id       string
			Password string
		}{}, types.UserNotExist
	}

	return struct {
		Id       string
		Password string
	}{Id: strconv.Itoa(int(user.Id)), Password: user.Password}, types.OK

}

// InsertUserDataToMySQL
//
//	@Description: 向数据库中插入用户数据并获取用户 Id
//	@receiver u
//	@param username
//	@param password
//	@param realName
//	@param identity
//	@return string
//	@return types.ErrNo
//	@data 2022-10-10 09:38:09
//	@author WuShaobei
func (u *UserDao) InsertUserDataToMySQL(username, password, realName string, identity types.IdentityTypes) (string, types.ErrNo) {
	var user types.ChineseCateringUser
	DB.Where("username = ?", username).Find(&user)
	if (user != types.ChineseCateringUser{}) {
		return "", types.UserHasExisted
	}

	user.Username = username
	user.Password = password
	user.RealName = realName
	user.Identity = identity

	DB.Create(&user)
	DB.Where("username = ?", username).Find(&user)

	if (user == types.ChineseCateringUser{}) {
		return "", types.UnknownError
	}
	return strconv.Itoa(int(user.Id)), types.OK
}

// SelectUsernameAndPasswordBySessionKeyFromRedis
//
//	@Description: 通过 SessionKey 从 Redis 获取用户名和密码
//	@receiver u
//	@param sessionKey
//	@return struct{
//	@return Password
//	@return types.ErrNo
//	@data 2022-10-10 09:46:08
//	@author WuShaobei
func (u *UserDao) SelectUsernameAndPasswordBySessionKeyFromRedis(sessionKey string) (struct{ Username, Password string }, types.ErrNo) {
	username, err := RDB.HGet(sessionKey, "username").Result()
	if err != nil {
		return struct{ Username, Password string }{}, types.UserNotExist
	}
	password, err := RDB.HGet(sessionKey, "password").Result()
	if err != nil {
		return struct{ Username, Password string }{}, types.UserNotExist
	}
	return struct{ Username, Password string }{Username: username, Password: password}, types.OK
}

// InsertUserDataToRedis
//
//	@Description: 向 Redis 写入用户数据并返回 SessionKey
//	@receiver u
//	@param username
//	@param password
//	@return string
//	@return types.ErrNo
//	@data 2022-10-10 09:59:31
//	@author WuShaobei
func (u *UserDao) InsertUserDataToRedis(username, password string) (string, types.ErrNo) {
	// 获取唯一标识符 uuid 作为该数据行的键
	sessionKey := uuid.NewV4().String()

	if err := RDB.HSet(sessionKey, "username", username).Err(); err != nil {
		log.Fatal(err)
		return "", types.UnknownError
	}
	if err := RDB.HSet(sessionKey, "password", password).Err(); err != nil {
		log.Fatal(err)
		return "", types.UnknownError
	}
	// 设置 3 天的存储时间
	_, err := RDB.Expire(sessionKey, 3*24*3600*time.Second).Result()
	if err != nil {
		return "", types.UnknownError
	} else {
		return sessionKey, types.OK
	}
}

func (u *UserDao) DeleteSessionKeyFromRedis(sessionKey string) types.ErrNo {
	RDB.Del(sessionKey)
	return types.OK
}
