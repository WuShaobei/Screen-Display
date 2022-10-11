package manager

import (
	"backEnd/dao"
	"backEnd/types"
)

type UserManager struct {
	//id       uint
	//username string
	//password string
	//realName string
	//identity string
}

func (u *UserManager) WhoAmI(id string) (types.WhoAmIData, types.ErrNo) {
	userDao := dao.UserDao{}
	return userDao.SelectByIdFromMySQL(id)
}

func (u *UserManager) LoginByPassword(username, password string, setCookie bool) (types.LoginData, types.ErrNo) {
	userDao := dao.UserDao{}

	data, errNo := userDao.SelectByUsernameFromMySQL(username)
	if errNo != types.OK {
		return types.LoginData{}, errNo
	}

	if password != data.Password {
		return types.LoginData{}, types.WrongPassword
	}

	if setCookie {
		sessionKey, errNo := userDao.InsertUserDataToRedis(username, password)
		if errNo != types.OK {
			return types.LoginData{}, errNo
		} else {
			return types.LoginData{Id: data.Id, SessionKey: sessionKey}, errNo
		}
	} else {
		return types.LoginData{Id: data.Id}, types.OK
	}
}

func (u *UserManager) LoginBySessionKey(sessionKey string) (types.LoginData, types.ErrNo) {
	userDao := dao.UserDao{}

	data, errNo := userDao.SelectBySessionKeyFromRedis(sessionKey)

	if errNo != types.OK {
		return types.LoginData{}, errNo
	}

	sqlData, errNo := userDao.SelectByUsernameFromMySQL(data.Username)
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

func (u *UserManager) Logout(SessionKey string) types.ErrNo {
	userDao := dao.UserDao{}
	return userDao.DeleteSessionKeyFromRedis(SessionKey)
}

func (u *UserManager) Register(body types.RegisterRequest) (string, types.ErrNo) {
	userDao := dao.UserDao{}
	return userDao.InsertUserDataToMySQL(body.Username, body.Password, body.RealName, body.Identity)
}
