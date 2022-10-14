// Package dao
// @Description: 用户管理 dao 类
package dao

import (
	"backEnd/types"
	"log"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserDao struct {
}

// SelectDataByIdFromMySQL
//
//	@Description: 通过 Id 从 mysql 获取用户信息
//	@receiver u *UserDao
//	@param id
//	@return types.WhoAmIData
//	@return types.ErrNo
func (u *UserDao) SelectDataByIdFromMySQL(id string) (types.WhoAmIData, types.ErrNo) {

	var user types.ChineseCateringUser
	DB.Where("id = ?", id).Find(&user)

	if (user != types.ChineseCateringUser{}) {
		return types.WhoAmIData{
			Id: strconv.Itoa(int(user.Id)), Username: user.Username, RealName: user.RealName, Identity: string(user.Identity),
		}, types.OK
	} else {
		return types.WhoAmIData{}, types.UserNotExist
	}
}

// SelectIdAndPasswordByUsernameFromMySQL
//
//	@Description: 通过 username 从 Mysql 读取 Id 和 password
//	@receiver u *UserDao
//	@param username
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
//	@receiver u *UserDao
//	@param username
//	@param password
//	@param realName
//	@param identity
//	@return string
//	@return types.ErrNo
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
//	@receiver u *UserDao
//	@param sessionKey
//	@return struct{ Username, Password string }
//	@return Password
//	@return types.ErrNo
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
//	@receiver u struct{ Username, Password string }
//	@param username
//	@param password
//	@return string
//	@return types.ErrNo
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

// DeleteSessionKeyFromRedis
//
//	@Description: 删除 redis 中 session 值为 sessionKey 的数据
//	@receiver u *UserDao
//	@param sessionKey
//	@return types.ErrNo
func (u *UserDao) DeleteSessionKeyFromRedis(sessionKey string) types.ErrNo {
	RDB.Del(sessionKey)
	return types.OK
}
