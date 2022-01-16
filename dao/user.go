package dao

import (
	"JD/modle"
	_ "github.com/go-sql-driver/mysql"
)

func ChangePassword(user modle.User) error {
	sqlStr := "update userBaseData set password = ? where username = ?"
	_, err := dB.Exec(sqlStr, user.Password, user.Password)
	if err != nil {
		return err
	}
	return err
}

func CheckPassword(user modle.User) (error, modle.User) {
	var check modle.User
	sqlStr := "select username,password from userBaseData where username = ?"
	err := dB.QueryRow(sqlStr, user.Username).Scan(&check.Username, &check.Password)
	if err != nil {
		return err, check
	}
	return err, check
}

func CheckUsername(user modle.User) error {
	var username string
	sqlStr := "select username from userBaseData where username = ?"
	err := dB.QueryRow(sqlStr, user.Username).Scan(&username)
	if err != nil {
		return err
	}
	return err
}

func WriteIn(user modle.User) error {
	sqlStr := "insert into userBaseData (username,password) values (?,?)"
	_, err := dB.Exec(sqlStr, user.Username, user.Password)
	if err != nil {
		return err
	}
	return err
}
