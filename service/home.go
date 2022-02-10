package service

import (
	"douban/dao"
	"douban/modle"
)

func Find(keyword string) (error, []int) {
	err, movieNums := dao.Find(keyword)
	if err != nil {
		return err, movieNums
	}
	return err, movieNums
}

func FindWithCategory(category string) (error, []modle.MovieInfo) {
	err, movies := dao.FindWithCategory(category)
	if err != nil {
		return err, movies
	}
	return err, movies
}
