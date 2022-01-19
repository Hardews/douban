package service

import (
	"douban/dao"
	"douban/modle"
)

func Find(keyword string) (error, []modle.MovieInfo) {
	err, movies := dao.Find(keyword)
	if err != nil {
		return err, movies
	}
	return err, movies
}
