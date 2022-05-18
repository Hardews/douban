package service

import (
	"douban/dao"
	"douban/model"
)

func Find(keyword string) (error, []model.MovieInfo) {
	err, movieNums := dao.Find(keyword)
	if err != nil {
		return err, movieNums
	}
	return err, movieNums
}

func FindWithCategory(category string) (error, []model.MovieInfo) {
	err, movies := dao.FindWithCategory(category)
	if err != nil {
		return err, movies
	}
	return err, movies
}
