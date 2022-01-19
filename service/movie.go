package service

import (
	"douban/dao"
	"douban/modle"
)

func GetComment(username string, num int) (error, []modle.UserComment) {
	err, comments := dao.GetComment(username, num)
	if err != nil {
		return err, comments
	}
	return err, comments
}

func GetMovieComment(username string, num int) (error, []modle.UserComment) {
	err, comments := dao.GetMovieComment(username, num)
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
