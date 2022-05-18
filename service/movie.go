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

func DeleteSeen(user model.UserSeen) error {
	return dao.DeleteSeen(user)
}

func DeleteWantSee(user model.UserWantSee) error {
	return dao.DeleteWantSee(user)
}

func SetCommentArea(area model.CommentArea) error {
	return dao.SetCommentArea(area)
}

func GetComment(num int) (error, []model.MovieReview) {
	return dao.GetComment(num)
}

func GetMovieComment(num int) (error, []model.ShortReview) {
	return dao.GetMovieComment(num)
}

func GetAMovieInfo(movieNum int) (error, model.MovieInfo) {
	return dao.GetAMovieInfo(movieNum)
}
