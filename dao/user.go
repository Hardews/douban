package dao

import (
	"douban/modle"
	_ "github.com/go-sql-driver/mysql"
)

func GetUserComment(username string) (error, []modle.UserComment, []modle.UserComment) {
	var shortComments, longComments []modle.UserComment
	sqlStr := "select movieNum,FilmCritics,time from shortComment where username = ?"
	rows, err := dB.Query(sqlStr, username)
	if err != nil {
		return err, shortComments, longComments
	}

	for rows.Next() {
		var shortComment modle.UserComment
		err := rows.Scan(&shortComment.MovieNum, &shortComment.Txt, &shortComment.Time)
		if err != nil {
			return err, shortComments, longComments
		}
		shortComments = append(shortComments, shortComment)
	}

	sqlStr = "select movieNum,Essay,time from longComment where username = ?"
	rows, err = dB.Query(sqlStr, username)
	if err != nil {
		return err, shortComments, longComments
	}
	defer rows.Close()

	for rows.Next() {
		var longComment modle.UserComment
		err := rows.Scan(&longComment.MovieNum, &longComment.Txt, &longComment.Time)
		if err != nil {
			return err, shortComments, longComments
		}
		longComments = append(longComments, longComment)
	}
	return err, shortComments, longComments
}

func UserWantSee(username, movieName string, movieNum int) error {
	sqlStr := "insert userWantSee (username,wantSee,num) values (?,?,?)"
	_, err := dB.Exec(sqlStr, username, movieName, movieNum)
	if err != nil {
		return err
	}
	return err
}

func SetIntroduce(username, introduce string) error {
	sqlStr := "update userMenu set introduce = ? where username = ?"
	_, err := dB.Exec(sqlStr, introduce, username)
	if err != nil {
		return err
	}
	return err
}

func UserMenuInfo(username string) (error, modle.UserInfoMenu) {
	var user modle.UserInfoMenu
	sqlStr := "select * from userMenu where username = ?"
	err := dB.QueryRow(sqlStr, username).Scan(&username, &user.Introduce, &user.FilmCritics, &user.Seen, &user.WantSee)
	if err != nil {
		return err, user
	}
	return err, user
}

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
	sqlStr = "insert into userMenu (username) values (?)"
	_, err = dB.Exec(sqlStr, user.Username)
	if err != nil {
		return err
	}
	return err
}
