package dao

import "douban/modle"

func GetComment(num int) (error, []modle.UserComment) {
	var comments []modle.UserComment
	sqlStr := "select Username,Essay,TIME from movieComment where movieNum = ?"
	rows, err := dB.Query(sqlStr, num)
	if err != nil {
		return err, comments
	}
	defer rows.Close()

	for rows.Next() {
		var comment modle.UserComment
		err = rows.Scan(&comment.Username, &comment.Txt, &comment.Time)
		if err != nil {
			return err, comments
		}
		comments = append(comments, comment)
	}
	return err, comments
}

func GetMovieComment(num int) (error, []modle.UserComment) {
	var comments []modle.UserComment
	sqlStr := "select Username,FilmCritics,time from shortComment where movieNum = ?"
	rows, err := dB.Query(sqlStr, num)
	if err != nil {
		return err, comments
	}
	defer rows.Close()

	for rows.Next() {
		var comment modle.UserComment
		err = rows.Scan(&comment.Username, &comment.Txt, &comment.Time)
		if err != nil {
			return err, comments
		}
		comments = append(comments, comment)
	}
	return err, comments
}

func Comment(Txt, username string, movieNum int) error {
	sqlStr := "insert shortComment (movieNum,Username,FilmCritics) values (?,?,?)"
	_, err := dB.Exec(sqlStr, movieNum, username, Txt)
	if err != nil {
		return err
	}
	return err
}

func CommentMovie(Txt, username string, movieNum int) error {
	sqlStr := "insert movieComment (movieNum,Username,Essay) values (?,?,?)"
	_, err := dB.Exec(sqlStr, movieNum, username, Txt)
	if err != nil {
		return err
	}
	return err
}

func FindWithCategory(category string) (error, []modle.MovieInfo) {
	var movies []modle.MovieInfo
	sqlStr := "select ChineseName,otherName,score,area,year,types,starring,director,commentNum,introduce,language,howLong,commentNum,seen,wantSee from movieBaseInfo where movieBaseInfo.types like ?"
	category = "%" + category + "%"
	rows, err := dB.Query(sqlStr, category)
	if err != nil {
		return err, movies
	}
	defer rows.Close()

	for rows.Next() {
		var movie modle.MovieInfo
		err := rows.Scan(&movie.Name, &movie.OtherName, &movie.Score, &movie.Area,
			&movie.Year, &movie.Types, &movie.Starring, &movie.Director, &movie.CommentNum, &movie.Introduce, &movie.Language,
			&movie.Time, &movie.CommentNum, &movie.Seen, &movie.WantSee)
		if err != nil {
			return err, movies
		}
		movies = append(movies, movie)
	}
	return err, movies
}

func GetAMovieInfo(movieNum int) (error, modle.MovieInfo) {
	var movie modle.MovieInfo
	sqlStr := "select ChineseName,otherName,score,area,year,types,starring,director,commentNum,introduce,language,howLong,commentNum,seen,wantSee from movieBaseInfo,movieExtraInfo where movieBaseInfo.num = ? and movieExtraInfo.num = ?"
	err := dB.QueryRow(sqlStr, movieNum, movieNum).Scan(&movie.Name, &movie.OtherName, &movie.Score, &movie.Area,
		&movie.Year, &movie.Types, &movie.Starring, &movie.Director, &movie.CommentNum, &movie.Introduce, &movie.Language,
		&movie.Time, &movie.CommentNum, &movie.Seen, &movie.WantSee)
	if err != nil {
		return err, movie
	}
	return err, movie
}
