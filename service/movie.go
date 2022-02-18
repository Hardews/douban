package service

import (
	"database/sql"
	"douban/dao"
	"douban/model"
)

func DeleteShortComment(username string, movieNum int) (error, bool) {
	err := dao.DeleteShortComment(username, movieNum)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, false
		}
		return err, false
	}
	return err, true
}

func DeleteLongComment(username string, movieNum int) (error, bool) {
	err := dao.DeleteLongComment(username, movieNum)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, false
		}
		return err, false
	}
	return err, true
}

func DeleteSeen(movieNum int, label, username string) error {
	err := dao.DeleteSeen(movieNum, label, username)
	if err != nil {
		return err
	}
	return err
}

func DeleteWantSee(movieNum int, label, username string) error {
	err := dao.DeleteWantSee(movieNum, label, username)
	if err != nil {
		return err
	}
	return err
}

func SetCommentArea(username, topic string, movieNum int) error {
	err := dao.SetCommentArea(username, topic, movieNum)
	if err != nil {
		return err
	}
	return err
}

func GetComment(num int) (error, []model.UserComment) {
	err, comments := dao.GetComment(num)
	if err != nil {
		return err, comments
	}
	return err, comments
}

func GetMovieComment(num int) (error, []model.UserComment) {
	err, comments := dao.GetMovieComment(num)
	if err != nil {
		return err, comments
	}
	return err, comments
}

func GetAMovieInfo(movieNum int) (error, model.MovieInfo) {
	err, movie := dao.GetAMovieInfo(movieNum)
	if err != nil {
		return err, movie
	}
	return err, movie
}
