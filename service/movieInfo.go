package service

import (
	"douban/dao"
	"douban/modle"
)

func GetAMovieInfo(movieNum int) (error, modle.MovieInfo) {
	err, movie := dao.GetAMovieInfo(movieNum)
	if err != nil {
		return err, movie
	}
	return err, movie
}
