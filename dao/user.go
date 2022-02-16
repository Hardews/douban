package dao

import (
	"database/sql"
	"douban/modle"
	_ "github.com/go-sql-driver/mysql"
)

func SetQuestion(username, question, answer string) (error, bool) {
	_, err := SelectQuestion(username)
	switch {
	case err != nil && err == sql.ErrNoRows:
		err = nil
	case err == nil:
		return err, false
	}

	sqlStr := "insert user_Encrypted (username,question,answer) values (?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, false
	}
	defer stmt.Close()

	_, err = stmt.Exec(sqlStr, username, question, answer)
	if err != nil {
		return err, false
	}
	return err, true
}

func CheckAnswer(username string) (error, string) {
	var answer string
	sqlStr := "select answer from user_Encrypted where username = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, ""
	}
	defer stmt.Close()
	err = stmt.QueryRow(sqlStr, username).Scan(&answer)
	if err != nil {
		return err, answer
	}
	return err, answer
}

func SelectQuestion(username string) (string, error) {
	var question string
	sqlStr := "select question from user_Encrypted where username = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	err = stmt.QueryRow(sqlStr, username).Scan(&question)
	if err != nil {
		return question, err
	}
	return question, err
}

func ChangePassword(user modle.User) error {
	sqlStr := "update user_Base_Data set password = ? where username = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sqlStr, user.Password, user.Username)
	if err != nil {
		return err
	}
	return err
}

func CheckPassword(user modle.User) (error, modle.User) {
	var check modle.User
	sqlStr := "select username,password from user_Base_Data where username = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, check
	}
	defer stmt.Close()

	err = stmt.QueryRow(sqlStr, user.Username).Scan(&check.Username, &check.Password)
	if err != nil {
		return err, check
	}
	return err, check
}

func CheckUsername(user modle.User) error {
	var username string
	sqlStr := "select username from user_Base_Data where username = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(sqlStr, user.Username).Scan(&username)
	if err != nil {
		return err
	}
	return err
}

func WriteIn(user modle.User) error {
	sqlStr := "insert into user_Base_Data (username,password,nickName) values (?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sqlStr, user.Username, user.Password, user.NickName)
	if err != nil {
		return err
	}

	sqlStr = "insert into user_Menu (username) values (?)"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(sqlStr, user.Username)
	if err != nil {
		return err
	}
	return err
}
