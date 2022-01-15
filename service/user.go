package service

import (
	"JD/dao"
	"JD/modle"
	"database/sql"
)

func CheckPassword(user modle.User) (error, bool) {
	err, check := dao.CheckPassword(user)
	if err != nil {
		return err, false
	}
	if check.Password != user.Password {
		return err, false
	}
	return err, true
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
