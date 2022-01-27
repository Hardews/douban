package service

import (
	"douban/dao"
	"douban/modle"
)

func SetCommentArea(username, topic string, movieNum int) error {
	err := dao.SetCommentArea(username, topic, movieNum)
	if err != nil {
		return err
	}
	return err
}

func GetComment(num int) (error, []modle.UserComment) {
	err, comments := dao.GetComment(num)
	if err != nil {
		return err, comments
	}
	return err, comments
}

func GetMovieComment(num int) (error, []modle.UserComment) {
	err, comments := dao.GetMovieComment(num)
	if err != nil {
		return err, comments
	}
	return err, comments
}

func FindWithCategory(category string) (error, []modle.MovieInfo) {
	err, movies := dao.FindWithCategory(category)
	if err != nil {
		return err, movies
	}
	return err, movies
}

func GetAMovieInfo(movieNum int) (error, modle.MovieInfo) {
	err, movie := dao.GetAMovieInfo(movieNum)
	if err != nil {
		return err, movie
	}
	return err, movie
}
