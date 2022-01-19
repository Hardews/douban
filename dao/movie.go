package dao

import "douban/modle"

func GetAMovieInfo(movieNum int) (error, modle.MovieInfo) {
	var movie modle.MovieInfo
	sqlStr := "select ChineseName,otherName,score,area,year,starring,director,commentNum,introduce,writer,language from movieBaseInfo,movieExtraInfo where movieBaseInfo.num = ? = movieExtraInfo.num"
	err := dB.QueryRow(sqlStr, movieNum).Scan(&movie.Name, &movie.OtherName, &movie.Score, &movie.Area,
		&movie.Year, &movie.Starring, &movie.Director, &movie.CommentNum, &movie.Introduce, &movie.Writer, &movie.Language)
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
