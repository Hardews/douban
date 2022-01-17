package service

import (
	"database/sql"
	"douban/dao"
	"douban/modle"
)

func SetIntroduce(username, introduce string) error {
	err := dao.SetIntroduce(username, introduce)
	if err != nil {
		return err
	}
	return err
}

func GetUserMenu(username string) (error, modle.UserInfoMenu) {
	err, user := dao.UserMenuInfo(username)
	if err != nil {
		return err, user
	}
	return err, user
}

func ChangePassword(user modle.User) error {
	var err error
	err, user.Password = Encryption(user.Password)
	if err != nil {
		return err
	}
	err = dao.ChangePassword(user)
	if err != nil {
		return err
	}
	return err
}

func CheckPassword(user modle.User) (error, bool) {
	err, check := dao.CheckPassword(user)
	if err != nil {
		return err, false
	}
	err, res := Interpretation(check.Password, user.Password)
	return err, res
}

func CheckUsername(user modle.User) (error, bool) {
	err := dao.CheckUsername(user)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, true
		}
		return err, false
	}
	return err, false
}

func WriteIn(user modle.User) error {
	err := dao.WriteIn(user)
	if err != nil {
		return err
	}
	return err
}
