package dao

import "douban/modle"

func Comment(Txt, username string, movieNum int) error {
	sqlStr := "insert movieComment (num,FUsername,FilmCritics) values (?,?,?)"
	_, err := dB.Exec(sqlStr, movieNum, username, Txt)
	if err != nil {
		return err
	}
	return err
}

func CommentMovie(Txt, username string, movieNum int) error {
	sqlStr := "insert movieComment (num,EUsername,Essay) values (?,?,?)"
	_, err := dB.Exec(sqlStr, movieNum, username, Txt)
	if err != nil {
		return err
	}
	return err
}

func FindWithCategory(category string) (error, []modle.MovieInfo) {
	var movies []modle.MovieInfo
	sqlStr := "select ChineseName,otherName,score,area,year,starring,director,types from movieBaseInfo where movieBaseInfo.types like ?"
	category = "%" + category + "%"
	rows, err := dB.Query(sqlStr, category)
	if err != nil {
		return err, movies
	}
	defer rows.Close()

	for rows.Next() {
		var movie modle.MovieInfo
		err := rows.Scan(&movie.Name, &movie.OtherName, &movie.Score, &movie.Area, &movie.Year, &movie.Starring, &movie.Director, &movie.Types)
		if err != nil {
			return err, movies
		}
		movies = append(movies, movie)
	}
	return err, movies
}

func GetAMovieInfo(movieNum int) (error, modle.MovieInfo) {
	var movie modle.MovieInfo
	sqlStr := "select ChineseName,otherName,score,area,year,types,starring,director,commentNum,introduce,writer,language from movieBaseInfo,movieExtraInfo where movieBaseInfo.num = ? and movieExtraInfo.num = ?"
	err := dB.QueryRow(sqlStr, movieNum, movieNum).Scan(&movie.Name, &movie.OtherName, &movie.Score, &movie.Area,
		&movie.Year, &movie.Types, &movie.Starring, &movie.Director, &movie.CommentNum, &movie.Introduce, &movie.Writer, &movie.Language)
	if err != nil {
		return err, movie
	}
	return err, movie
}

func FindMovieNumByName(movieName string) (error, int) {
	var movieNum int
	sqlStr := "select num from movieBaseInfo where chineseName = ?"
	err := dB.QueryRow(sqlStr, movieName).Scan(&movieNum)
	if err != nil {
		return err, movieNum
	}
	return err, movieNum
}
