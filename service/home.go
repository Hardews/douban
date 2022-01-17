package service

import (
	"douban/dao"
	"douban/modle"
)

func GetMovieBaseInfo() (error, []modle.MovieInfo) {
	err, infos := dao.GetMovieBaseInfo()
	if err != nil {
		return err, infos
	}
	return err, infos
}

func GetMovieAllInfo() (error, []modle.MovieInfo) {
	err, infos := dao.GetMovieAllInfo()
	if err != nil {
		return err, infos
	}
	return err, infos
}
