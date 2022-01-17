package service

import (
	"douban/dao"
	"douban/modle"
)

func GetMovieBaseInfo() (error, []modle.MovieBaseInfo) {
	err, infos := dao.GetMovieBaseInfo()
	if err != nil {
		return err, infos
	}
	return err, infos
}
