package manager

import (
	"backEnd/dao"
	"backEnd/types"
)

type UserManager struct {
	userDao dao.UserDao
}

func (u *UserManager) WhoAmI(id string) (types.WhoAmIData, types.ErrNo) {
	return u.userDao.SelectDataByIdFromMySQL(id)
}

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

func (u *UserManager) Logout(SessionKey string) types.ErrNo {
	return u.userDao.DeleteSessionKeyFromRedis(SessionKey)
}

func (u *UserManager) Register(body types.RegisterRequest) (string, types.ErrNo) {
	return u.userDao.InsertUserDataToMySQL(body.Username, body.Password, body.RealName, body.Identity)
}
