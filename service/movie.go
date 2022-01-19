package service

import (
	"douban/dao"
	"douban/modle"
)

func FindMovieNumByName(movieName string) (error, int) {
	err, movieNum := dao.FindMovieNumByName(movieName)
	if err != nil {
		return err, movieNum
	}
	return err, movieNum
}

func GetAMovieInfo(movieNum int) (error, modle.MovieInfo) {
	err, movie := dao.GetAMovieInfo(movieNum)
	if err != nil {
		return err, movie
	}
	return err, movie
}
