// Package manager
// @Description: 用户 Manager 类

package manager

import (
	"backEnd/dao"
	"backEnd/types"
)

type UserManager struct {
	userDao dao.UserDao
}

// WhoAmI
//
//	@Description: 获取用户数据
//	@receiver u *UserManager
//	@param id
//	@return types.WhoAmIData
//	@return types.ErrNo
func (u *UserManager) WhoAmI(id string) (types.WhoAmIData, types.ErrNo) {
	return u.userDao.SelectDataByIdFromMySQL(id)
}

// LoginByPassword
//
//	@Description: 通过密码登录
//	@receiver u *UserManager
//	@param username
//	@param password
//	@param setCookie
//	@return types.LoginData
//	@return types.ErrNo9
//	@author WuShaobei
func (u *UserManager) LoginByPassword(username, password string, setCookie bool) (types.LoginData, types.ErrNo) {

	data, errNo := u.userDao.SelectIdAndPasswordByUsernameFromMySQL(username)
	if errNo != types.OK {
		return types.LoginData{}, errNo
	}

	if password != data.Password {
		return types.LoginData{}, types.WrongPassword
	}

	if setCookie {
		sessionKey, errNo := u.userDao.InsertUserDataToRedis(username, password)
		if errNo != types.OK {
			return types.LoginData{}, errNo
		} else {
			return types.LoginData{Id: data.Id, SessionKey: sessionKey}, errNo
		}
	} else {
		return types.LoginData{Id: data.Id}, types.OK
	}
}

// LoginBySessionKey
//
//	@Description: 通过 session 值登录
//	@receiver u *UserManager
//	@param sessionKey
//	@return types.LoginData
//	@return types.ErrNo
func (u *UserManager) LoginBySessionKey(sessionKey string) (types.LoginData, types.ErrNo) {

	data, errNo := u.userDao.SelectUsernameAndPasswordBySessionKeyFromRedis(sessionKey)

	if errNo != types.OK {
		return types.LoginData{}, errNo
	}

	sqlData, errNo := u.userDao.SelectIdAndPasswordByUsernameFromMySQL(data.Username)
	if errNo != types.OK {
		return types.LoginData{}, errNo
	}

	if sqlData.Password == data.Password {
		return types.LoginData{
			Id: sqlData.Id,
		}, types.OK
	} else {
		return types.LoginData{}, types.WrongPassword
	}
}

// Logout
//
//	@Description: 登出
//	@receiver u *UserManager
//	@param SessionKey
//	@return types.ErrNo
func (u *UserManager) Logout(SessionKey string) types.ErrNo {
	return u.userDao.DeleteSessionKeyFromRedis(SessionKey)
}

// Register
//
//	@Description: 注册
//	@receiver u *UserManager
//	@param body
//	@return string
//	@return types.ErrNo
func (u *UserManager) Register(body types.RegisterRequest) (string, types.ErrNo) {
	return u.userDao.InsertUserDataToMySQL(body.Username, body.Password, body.RealName, body.Identity)
}
