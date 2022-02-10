package service

import (
	"douban/dao"
	"douban/modle"
)

func NewMovie(movie modle.MovieInfo) (error, int) {
	err, newNum := dao.NewMovie(movie)
	if err != nil {
		return err, newNum
	}
	return err, newNum
}
