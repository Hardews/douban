package dao

import (
	"douban/modle"
)

func NewMovie(movie modle.MovieInfo) (error, int) {
	sqlStr := "insert movie_Base_Info (chineseName,otherName,score,area,year,starring,director,types) values (?,?,?,?,?,?,?,?)"
	_, err := dB.Exec(sqlStr, movie.Name, movie.OtherName, movie.Score, movie.Area, movie.Year, movie.Starring, movie.Director, movie.Types)
	if err != nil {
		return err, 0
	}

	var movieNum int
	sqlStr = "SELECT num from movie_Base_Info where num = (SELECT max(num) FROM movieBaseInfo);"
	err = dB.QueryRow(sqlStr).Scan(&movieNum)
	if err != nil {
		return err, 0
	}

	sqlStr = "insert movie_Extra_Info (num,movieName,introduce,howLong) values (?,?,?,?) "
	_, err = dB.Exec(sqlStr, movieNum, movie.Name, movie.Introduce, movie.Time)
	if err != nil {
		return err, 0
	}

	return err, movieNum
}
