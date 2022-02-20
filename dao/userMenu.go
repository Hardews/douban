package dao

import (
	"douban/model"
	"strconv"
)

func GetWantSee(username string) (error, []model.UserHistory) {
	var wantSees []model.UserHistory
	sqlStr := "select comment,num,label from user_Want_See where username = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, wantSees
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		return err, wantSees
	}
	defer rows.Close()

	sqlStr = "select img from movie_Base_Info where num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err, wantSees
	}

	for rows.Next() {
		var wantSee model.UserHistory
		err := rows.Scan(&wantSee.Comment, &wantSee.MovieNum, &wantSee.Label)
		if err != nil {
			return err, wantSees
		}

		err = stmt.QueryRow(wantSee.MovieNum).Scan(&wantSee.Img)
		if err != nil {
			return err, wantSees
		}

		movieNum := strconv.Itoa(wantSee.MovieNum)
		wantSee.Url = "http://49.235.99.195:8090/movieInfo/" + movieNum
		wantSees = append(wantSees, wantSee)
	}
	return err, wantSees
}

func GetSeen(username string) (error, []model.UserHistory) {
	var Seens []model.UserHistory
	sqlStr := "select comment,num,label from user_Seen where username = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, Seens
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		return err, Seens
	}
	defer rows.Close()

	sqlStr = "select img from movie_Base_Info where num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err, Seens
	}

	for rows.Next() {
		var Seen model.UserHistory
		err := rows.Scan(&Seen.Comment, &Seen.MovieNum, &Seen.Label)
		if err != nil {
			return err, Seens
		}
		err = stmt.QueryRow(Seen.MovieNum).Scan(&Seen.Img)

		movieNum := strconv.Itoa(Seen.MovieNum)
		Seen.Url = "http://49.235.99.195:8090/movieInfo/" + movieNum
		Seens = append(Seens, Seen)
	}
	return err, Seens
}

func UserSeen(username, comment, label string, movieNum int) error {
	sqlStr := "insert user_Seen (username,comment,num,label) values (?,?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, comment, movieNum, label)
	if err != nil {
		return err
	}

	sqlStr = "select Seen from movie_Extra_Info where num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	var num int
	err = stmt.QueryRow(movieNum).Scan(&num)
	if err != nil {
		return err
	}

	num += 1
	sqlStr = "update movie_Extra_Info set Seen = ? where num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(num, movieNum)
	if err != nil {
		return err
	}
	return err
}

func UserWantSee(username, comment, label string, movieNum int) error {
	sqlStr := "insert user_Want_See (username,comment,num,label) values (?,?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, comment, movieNum, label)
	if err != nil {
		return err
	}

	sqlStr = "select wantSee from movie_Extra_Info where num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	var num int
	err = stmt.QueryRow(movieNum).Scan(&num)
	if err != nil {
		return err
	}

	num += 1
	sqlStr = "update movie_Extra_Info set wantSee = ? where num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(num, movieNum)
	if err != nil {
		return err
	}
	return err
}

func GetUserComment(username string) (error, []model.UserComment, []model.UserComment) {
	var shortComments, longComments []model.UserComment
	sqlStr := "select movieNum,FilmCritics,time from short_Comment where username = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, shortComments, longComments
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		return err, shortComments, longComments
	}

	for rows.Next() {
		var shortComment model.UserComment
		err := rows.Scan(&shortComment.MovieNum, &shortComment.Txt, &shortComment.Time)
		if err != nil {
			return err, shortComments, longComments
		}
		movieNum := strconv.Itoa(shortComment.MovieNum)
		shortComment.Url = "http://49.235.99.195:8090/movieInfo/" + movieNum
		shortComments = append(shortComments, shortComment)
	}

	sqlStr = "select movieNum,Essay,time from long_Comment where username = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err, shortComments, longComments
	}
	rows, err = stmt.Query(username)
	if err != nil {
		return err, shortComments, longComments
	}
	defer rows.Close()

	for rows.Next() {
		var longComment model.UserComment
		err := rows.Scan(&longComment.MovieNum, &longComment.Txt, &longComment.Time)
		if err != nil {
			return err, shortComments, longComments
		}
		movieNum := strconv.Itoa(longComment.MovieNum)
		longComment.Url = "http://49.235.99.195:8090/movieInfo/" + movieNum
		longComments = append(longComments, longComment)
	}
	return err, shortComments, longComments
}

func SetIntroduce(username, introduce string) error {
	sqlStr := "update user_Menu set introduce = ? where username = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(introduce, username)
	if err != nil {
		return err
	}
	return err
}

func UserMenuInfo(username string) (error, model.User) {
	var user model.User
	sqlStr := "select username,introduce from user_Menu where username = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, user
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&username, &user.Introduce)
	if err != nil {
		return err, user
	}

	sqlStr = "select nickName,avatar,address from user_Base_Data where username = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err, user
	}
	err = stmt.QueryRow(username).Scan(&user.NickName, &user.Img, &user.ImgAddress)
	if err != nil {
		return err, user
	}

	return err, user
}
