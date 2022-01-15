package dao

import "JD/modle"

func CheckUsername(user modle.User) error {
	var username string
	sqlStr := "select username from userBaseData where username = ?"
	err := dB.QueryRow(sqlStr, user.Username).Scan(username)
	if err != nil {
		return err
	}
	return err
}

func WriteIn(user modle.User) error {
	sqlStr := "insert into user (username,password) values (?,?)"
	_, err := dB.Exec(sqlStr, user.Username, user.Password)
	if err != nil {
		return err
	}
	return err
}
