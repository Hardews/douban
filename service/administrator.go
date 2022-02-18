package service

import (
	"douban/dao"
	"douban/model"
)

func NewMovie(movie model.MovieInfo) (error, int) {
	err, newNum := dao.NewMovie(movie)
	if err != nil {
		return err, newNum
	}
	return err, newNum
}
