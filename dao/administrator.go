package dao

import (
	"douban/model"
)

func NewMovie(movie model.MovieInfo) (error, int) {
	sqlStr := "insert movie_Base_Info (chineseName,otherName,score,area,year,starring,director,types,img,address) values (?,?,?,?,?,?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, 0
	}
	defer stmt.Close()

	_, err = stmt.Exec(movie.Name, movie.OtherName, movie.Score, movie.Area, movie.Year, movie.Starring, movie.Director, movie.Types, movie.Img, movie.ImgAddress)
	if err != nil {
		return err, 0
	}

	var movieNum int
	sqlStr = "SELECT num from movie_Base_Info where num = (SELECT max(num) FROM movieBaseInfo);"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err, 0
	}

	err = stmt.QueryRow().Scan(&movieNum)
	if err != nil {
		return err, 0
	}

	sqlStr = "insert movie_Extra_Info (num,movieName,introduce,howLong) values (?,?,?,?) "
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err, 0
	}

	_, err = stmt.Exec(movieNum, movie.Name, movie.Introduce, movie.Time)
	if err != nil {
		return err, 0
	}

	return err, movieNum
}
