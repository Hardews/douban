package dao

import (
	"douban/modle"
	"strconv"
)

func GetWantSee(username string) (error, []modle.UserHistory) {
	var wantSees []modle.UserHistory
	sqlStr := "select comment,num,label from userWantSee where username = ?"
	rows, err := dB.Query(sqlStr, username)
	if err != nil {
		return err, wantSees
	}
	defer rows.Close()

	for rows.Next() {
		var wantSee modle.UserHistory
		err := rows.Scan(&wantSee.Comment, &wantSee.MovieNum, &wantSee.Label)
		if err != nil {
			return err, wantSees
		}
		movieNum := strconv.Itoa(wantSee.MovieNum)
		wantSee.Url = "http://101.201.234.29:8080/movieInfo/" + movieNum
		wantSees = append(wantSees, wantSee)
	}
	return err, wantSees
}

func GetSeen(username string) (error, []modle.UserHistory) {
	var Seens []modle.UserHistory
	sqlStr := "select comment,num,label from userSeen where username = ?"
	rows, err := dB.Query(sqlStr, username)
	if err != nil {
		return err, Seens
	}
	defer rows.Close()

	for rows.Next() {
		var Seen modle.UserHistory
		err := rows.Scan(&Seen.Comment, &Seen.MovieNum, &Seen.Label)
		if err != nil {
			return err, Seens
		}
		movieNum := strconv.Itoa(Seen.MovieNum)
		Seen.Url = "http://101.201.234.29:8080/movieInfo/" + movieNum
		Seens = append(Seens, Seen)
	}
	return err, Seens
}

func UserSeen(username, comment, label string, movieNum int) error {
	sqlStr := "insert userSeen (username,comment,num,label) values (?,?,?,?)"
	_, err := dB.Exec(sqlStr, username, comment, movieNum, label)
	if err != nil {
		return err
	}

	sqlStr = "select Seen from movieBaseInfo where num = ?"
	var num int
	err = dB.QueryRow(sqlStr, movieNum).Scan(&num)
	if err != nil {
		return err
	}

	num += 1
	sqlStr = "update movieBaseInfo set Seen = ? where num = ?"
	_, err = dB.Exec(sqlStr, num, movieNum)
	if err != nil {
		return err
	}
	return err
}

func UserWantSee(username, comment, label string, movieNum int) error {
	sqlStr := "insert userWantSee (username,comment,num,label) values (?,?,?,?)"
	_, err := dB.Exec(sqlStr, username, comment, movieNum, label)
	if err != nil {
		return err
	}

	sqlStr = "select wantSee from movieBaseInfo where num = ?"
	var num int
	err = dB.QueryRow(sqlStr, movieNum).Scan(&num)
	if err != nil {
		return err
	}

	num += 1
	sqlStr = "update movieBaseInfo set wantSee = ? where num = ?"
	_, err = dB.Exec(sqlStr, num, movieNum)
	if err != nil {
		return err
	}
	return err
}

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
		movieNum := strconv.Itoa(shortComment.MovieNum)
		shortComment.Url = "http://101.201.234.29:8080/movieInfo/" + movieNum
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
		movieNum := strconv.Itoa(longComment.MovieNum)
		longComment.Url = "http://101.201.234.29:8080/movieInfo/" + movieNum
		longComments = append(longComments, longComment)
	}
	return err, shortComments, longComments
}

func SetIntroduce(username, introduce string) error {
	sqlStr := "update userMenu set introduce = ? where username = ?"
	_, err := dB.Exec(sqlStr, introduce, username)
	if err != nil {
		return err
	}
	return err
}

func UserMenuInfo(username string) (error, modle.User) {
	var user modle.User
	sqlStr := "select username,introduce from userMenu where username = ?"
	err := dB.QueryRow(sqlStr, username).Scan(&username, &user.Introduce)
	if err != nil {
		return err, user
	}

	sqlStr = "select nickName from userBaseData where username = ?"
	err = dB.QueryRow(sqlStr, username).Scan(&user.NickName)
	if err != nil {
		return err, user
	}

	return err, user
}
