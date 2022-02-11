package dao

import (
	"database/sql"
	"douban/modle"
	"strconv"
)

func DeleteShortComment(username string, movieNum int) error {
	var iUsername string
	sqlStr := "select FilmCritics from short_Comment where username = ? and movieNum = ?"
	err := dB.QueryRow(sqlStr, username, movieNum).Scan(&iUsername)
	if err != nil {
		return err
	}

	username = username + "已删除"
	sqlStr = "update short_Comment set username = ? where username = ? and movieNum = ?"
	_, err = dB.Exec(sqlStr, username, iUsername, movieNum)
	if err != nil {
		return err
	}
	return err
}

func DeleteLongComment(username string, movieNum int) error {
	var iUsername string
	sqlStr := "select essay from movie_Comment where username = ? and movieNum = ?"
	err := dB.QueryRow(sqlStr, username, movieNum).Scan(&iUsername)
	if err != nil {
		return err
	}

	username = username + "已删除"
	sqlStr = "update essay set username = ? where username = ? and movieNum = ?"
	_, err = dB.Exec(sqlStr, username, iUsername, movieNum)
	if err != nil {
		return err
	}

	sqlStr = "select commentNum from movie_Extra_Info where num = ?"
	var num int
	err = dB.QueryRow(sqlStr, movieNum).Scan(&num)
	if err != nil {
		return err
	}

	num -= 1
	sqlStr = "update movie_Extra_Info set commentNum = ? where num = ?"
	_, err = dB.Exec(sqlStr, num, movieNum)
	if err != nil {
		return err
	}

	return err
}

func DeleteSeen(movieNum int, label, username string) error {
	username = username + "已删除"
	sqlStr := "update userSeen set username = ? where movieNum = ? and label = ?"
	_, err := dB.Exec(sqlStr, username, movieNum, label)
	if err != nil {
		return err
	}

	sqlStr = "select seen from movie_Base_Info where num = ?"
	var num int
	err = dB.QueryRow(sqlStr, movieNum).Scan(&num)
	if err != nil {
		return err
	}

	num -= 1
	sqlStr = "update movie_Base_Info set wantSee = ? where num = ?"
	_, err = dB.Exec(sqlStr, num, movieNum)
	if err != nil {
		return err
	}
	return err
}

func DeleteWantSee(movieNum int, label, username string) error {
	username = username + "已删除"
	sqlStr := "update user_Want_See set username = ? where movieNum = ? and label = ?"
	_, err := dB.Exec(sqlStr, username, movieNum, label)
	if err != nil {
		return err
	}

	sqlStr = "select wantSee from movie_Base_Info where num = ?"
	var num int
	err = dB.QueryRow(sqlStr, movieNum).Scan(&num)
	if err != nil {
		return err
	}

	num -= 1
	sqlStr = "update movie_Base_Info set wantSee = ? where num = ?"
	_, err = dB.Exec(sqlStr, num, movieNum)
	if err != nil {
		return err
	}
	return err
}

func GetComment(num int) (error, []modle.UserComment) {
	var comments []modle.UserComment
	sqlStr := "select Username,Essay,TIME,commentTopic from movie_Comment where movieNum = ?"
	rows, err := dB.Query(sqlStr, num)
	if err != nil {
		return err, comments
	}
	defer rows.Close()

	for rows.Next() {
		var comment modle.UserComment
		err = rows.Scan(&comment.Username, &comment.Txt, &comment.Time, &comment.Topic)
		if err != nil {
			return err, comments
		}
		movieNum := strconv.Itoa(num)
		comment.Url = "http://101.201.234.29:8080/movieInfo/" + movieNum
		comments = append(comments, comment)
	}
	return err, comments
}

func GetMovieComment(num int) (error, []modle.UserComment) {
	var comments []modle.UserComment
	sqlStr := "select Username,FilmCritics,time from short_Comment where movieNum = ?"
	rows, err := dB.Query(sqlStr, num)
	if err != nil {
		return err, comments
	}
	defer rows.Close()

	for rows.Next() {
		var comment modle.UserComment
		comment.MovieNum = num
		err = rows.Scan(&comment.Username, &comment.Txt, &comment.Time)
		if err != nil {
			return err, comments
		}
		movieNum := strconv.Itoa(num)
		comment.Url = "http://101.201.234.29:8080/movieInfo/" + movieNum
		comments = append(comments, comment)
	}
	return err, comments
}

func UpdateShortComment(username, txt string, movieNum int) error {
	sqlStr := "update short_Comment set FilmCritics = ? where username = ? and movieNum = ?"
	_, err := dB.Exec(sqlStr, txt, username, movieNum)
	if err != nil {
		return err
	}
	return err
}

func SelectShortComment(username string, movieNum int) (error, bool) {
	var iTxt string
	sqlStr := "select FilmCritics from short_Comment where username = ? and movieNum = ?"
	err := dB.QueryRow(sqlStr, username, movieNum).Scan(&iTxt)
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
	_, err := dB.Exec(sqlStr, movieNum, username, Txt)
	if err != nil {
		return err
	}
	return err
}

func UpdateLongComment(username, txt string, movieNum int) error {
	sqlStr := "update movie_Comment set Essay = ? where username = ? and movieNum = ?"
	_, err := dB.Exec(sqlStr, txt, username, movieNum)
	if err != nil {
		return err
	}
	return err
}

func SelectLongComment(username string, movieNum int) (error, bool) {
	var iTxt string
	sqlStr := "select Essay from movie_Comment where username = ? and movieNum = ?"
	err := dB.QueryRow(sqlStr, username, movieNum).Scan(&iTxt)
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
	_, err := dB.Exec(sqlStr, movieNum, username, Txt, commentTopic)
	if err != nil {
		return err
	}
	return err
}

func FindWithCategory(category string) (error, []modle.MovieInfo) {
	var movies []modle.MovieInfo
	sqlStr := "select num,ChineseName,otherName,score,area,year,types,starring,director from movie_Base_Info where types like ?"
	category = "%" + category + "%"
	rows, err := dB.Query(sqlStr, category)
	if err != nil {
		return err, movies
	}
	defer rows.Close()

	for rows.Next() {
		var movie modle.MovieInfo
		err := rows.Scan(&movie.Num, &movie.Name, &movie.OtherName, &movie.Score, &movie.Area,
			&movie.Year, &movie.Types, &movie.Starring, &movie.Director)
		if err != nil {
			return err, movies
		}
		movieNum := strconv.Itoa(movie.Num)
		movie.Url = "http://101.201.234.29:8080/movieInfo/" + movieNum
		movies = append(movies, movie)
	}
	return err, movies
}

func GetAMovieInfo(movieNum int) (error, modle.MovieInfo) {
	var movie modle.MovieInfo
	sqlStr := "select ChineseName,otherName,score,area,year,types,starring,director,commentNum,introduce,howLong,commentNum,seen,wantSee,img from movie_Base_Info,movie_Extra_Info where movie_Base_Info.num = ? and movie_Extra_Info.num = ?"
	err := dB.QueryRow(sqlStr, movieNum, movieNum).Scan(&movie.Name, &movie.OtherName, &movie.Score, &movie.Area,
		&movie.Year, &movie.Types, &movie.Starring, &movie.Director, &movie.CommentNum, &movie.Introduce,
		&movie.Time, &movie.CommentNum, &movie.Seen, &movie.WantSee, &movie.Img)
	Num := strconv.Itoa(movieNum)
	movie.Num = movieNum
	movie.Url = "http://101.201.234.29:8080/movieInfo/" + Num
	if err != nil {
		return err, movie
	}
	return err, movie
}
