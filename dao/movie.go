package dao

import (
	"database/sql"
	"douban/model"
	"strconv"
)

func DeleteShortComment(username string, movieNum int) error {
	var iUsername string
	sqlStr := "select FilmCritics from short_Comment where username = ? and movieNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, movieNum).Scan(&iUsername)
	if err != nil {
		return err
	}

	username = username + "已删除"
	sqlStr = "update short_Comment set username = ? where username = ? and movieNum = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username, iUsername, movieNum)
	if err != nil {
		return err
	}
	return err
}

func DeleteLongComment(username string, movieNum int) error {
	var iUsername string
	sqlStr := "select essay from movie_Comment where username = ? and movieNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, movieNum).Scan(&iUsername)
	if err != nil {
		return err
	}

	username = username + "已删除"
	sqlStr = "update essay set username = ? where username = ? and movieNum = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username, iUsername, movieNum)
	if err != nil {
		return err
	}

	sqlStr = "select commentNum from movie_Extra_Info where num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	var num int
	err = stmt.QueryRow(movieNum).Scan(&num)
	if err != nil {
		return err
	}

	num -= 1
	sqlStr = "update movie_Extra_Info set commentNum = ? where num = ?"
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

func DeleteSeen(movieNum int, label, username string) error {
	username = username + "已删除"
	sqlStr := "update userSeen set username = ? where movieNum = ? and label = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, movieNum, label)
	if err != nil {
		return err
	}

	sqlStr = "select seen from movie_Base_Info where num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	var num int
	err = stmt.QueryRow(movieNum).Scan(&num)
	if err != nil {
		return err
	}

	num -= 1
	sqlStr = "update movie_Base_Info set wantSee = ? where num = ?"
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

func DeleteWantSee(movieNum int, label, username string) error {
	username = username + "已删除"
	sqlStr := "update user_Want_See set username = ? where movieNum = ? and label = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, movieNum, label)
	if err != nil {
		return err
	}

	sqlStr = "select wantSee from movie_Base_Info where num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	var num int
	err = stmt.QueryRow(movieNum).Scan(&num)
	if err != nil {
		return err
	}

	num -= 1
	sqlStr = "update movie_Base_Info set wantSee = ? where num = ?"
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

func GetComment(num int) (error, []model.UserComment) {
	var comments []model.UserComment
	sqlStr := "select Username,Essay,TIME,commentTopic from movie_Comment where movieNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, comments
	}
	defer stmt.Close()

	rows, err := stmt.Query(num)
	if err != nil {
		return err, comments
	}
	defer rows.Close()

	for rows.Next() {
		var comment model.UserComment
		err = rows.Scan(&comment.Username, &comment.Txt, &comment.Time, &comment.Topic)
		if err != nil {
			return err, comments
		}
		movieNum := strconv.Itoa(num)
		comment.Url = "http://49.235.99.195:8080/movieInfo/" + movieNum
		comments = append(comments, comment)
	}
	return err, comments
}

func GetMovieComment(num int) (error, []model.UserComment) {
	var comments []model.UserComment
	sqlStr := "select Username,FilmCritics,time from short_Comment where movieNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, comments
	}
	defer stmt.Close()

	rows, err := stmt.Query(num)
	if err != nil {
		return err, comments
	}
	defer rows.Close()

	for rows.Next() {
		var comment model.UserComment
		comment.MovieNum = num
		err = rows.Scan(&comment.Username, &comment.Txt, &comment.Time)
		if err != nil {
			return err, comments
		}
		movieNum := strconv.Itoa(num)
		comment.Url = "http://49.235.99.195:8080/movieInfo/" + movieNum
		comments = append(comments, comment)
	}
	return err, comments
}

func UpdateShortComment(username, txt string, movieNum int) error {
	sqlStr := "update short_Comment set FilmCritics = ? where username = ? and movieNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(txt, username, movieNum)
	if err != nil {
		return err
	}
	return err
}

func SelectShortComment(username string, movieNum int) (error, bool) {
	var iTxt string
	sqlStr := "select FilmCritics from short_Comment where username = ? and movieNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, false
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, movieNum).Scan(&iTxt)
	switch {
	case err == nil:
		return err, false
	case err != nil && err == sql.ErrNoRows:
		return nil, true
	default:
		return err, false
	}
}

func Comment(Txt, username string, movieNum int) error {
	sqlStr := "insert short_Comment (movieNum,Username,FilmCritics) values (?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(movieNum, username, Txt)
	if err != nil {
		return err
	}
	return err
}

func UpdateLongComment(username, txt string, movieNum int) error {
	sqlStr := "update movie_Comment set Essay = ? where username = ? and movieNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(txt, username, movieNum)
	if err != nil {
		return err
	}
	return err
}

func SelectLongComment(username string, movieNum int) (error, bool) {
	var iTxt string
	sqlStr := "select Essay from movie_Comment where username = ? and movieNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, false
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, movieNum).Scan(&iTxt)
	switch {
	case err == nil:
		return err, false
	case err != nil && err == sql.ErrNoRows:
		return nil, true
	default:
		return err, false
	}
}

func CommentMovie(Txt, username, commentTopic string, movieNum int) error {
	sqlStr := "insert movie_Comment (movieNum,Username,Essay,commentTopic) values (?,?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(movieNum, username, Txt, commentTopic)
	if err != nil {
		return err
	}
	return err
}

func FindWithCategory(category string) (error, []model.MovieInfo) {
	var movies []model.MovieInfo
	sqlStr := "select num,ChineseName,otherName,score,area,year,types,starring,director from movie_Base_Info where types like ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, movies
	}
	defer stmt.Close()

	category = "%" + category + "%"
	rows, err := stmt.Query(category)
	if err != nil {
		return err, movies
	}
	defer rows.Close()

	for rows.Next() {
		var movie model.MovieInfo
		err := rows.Scan(&movie.Num, &movie.Name, &movie.OtherName, &movie.Score, &movie.Area,
			&movie.Year, &movie.Types, &movie.Starring, &movie.Director)
		if err != nil {
			return err, movies
		}
		movieNum := strconv.Itoa(movie.Num)
		movie.Url = "http://49.235.99.195:8080/movieInfo/" + movieNum
		movies = append(movies, movie)
	}
	return err, movies
}

func GetAMovieInfo(movieNum int) (error, model.MovieInfo) {
	var movie model.MovieInfo
	sqlStr := "select ChineseName,otherName,score,area,year,types,starring,director,commentNum,introduce,howLong,commentNum,seen,wantSee,img,address from movie_Base_Info,movie_Extra_Info where movie_Base_Info.num = ? and movie_Extra_Info.num = ?"

	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, movie
	}
	defer stmt.Close()

	err = stmt.QueryRow(movieNum, movieNum).Scan(&movie.Name, &movie.OtherName, &movie.Score, &movie.Area,
		&movie.Year, &movie.Types, &movie.Starring, &movie.Director, &movie.CommentNum, &movie.Introduce,
		&movie.Time, &movie.CommentNum, &movie.Seen, &movie.WantSee, &movie.Img, &movie.ImgAddress)
	Num := strconv.Itoa(movieNum)
	movie.Num = movieNum
	movie.Url = "http://49.235.99.195:8080/movieInfo/" + Num
	if err != nil {
		return err, movie
	}
	return err, movie
}
