package service

import (
	"database/sql"
	"douban/dao"
	"douban/model"
)

func DoNotLikeTopic(username string, areaNum int) error {
	err := dao.DoNotLikeTopic(username, areaNum)
	if err != nil {
		return err
	}
	return err
}

func DoNotLikeComment(username string, areaNum, commentNum int) error {
	err := dao.DoNotLikeComment(username, areaNum, commentNum)
	if err != nil {
		return err
	}
	return err
}

func DeleteComment(username string, movieNum, areaNum int) (error, bool) {
	err := dao.DeleteComment(username, movieNum, areaNum)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, false
		}
		return err, false
	}
	return err, true
}

func DeleteCommentArea(movieNum, areaNum int) (error, bool) {
	err := dao.DeleteCommentArea(movieNum, areaNum)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, false
		}
		return err, false
	}
	return err, true
}

func GiveCommentLike(username, name string, movieNum, areaNum int) (error, bool) {
	err, flag := dao.GiveCommentLike(username, name, movieNum, areaNum)
	if err != nil {
		return err, flag
	}
	return err, flag
}

func GiveTopicLike(username string, movieNum, num int) (error, bool) {
	err, flag := dao.GiveTopicLike(username, movieNum, num)
	if err != nil {
		return err, flag
	}
	return err, flag
}

func GetCommentArea(movieNum int) (error, []model.CommentArea) {
	err, commentAreas := dao.GetCommentArea(movieNum)
	if err != nil {
		return err, commentAreas
	}
	return err, commentAreas
}

func GetCommentByNum(movieNum, areaNum int) (error, []model.CommentArea) {
	err, comments := dao.GetCommentByNum(movieNum, areaNum)
	if err != nil {
		return err, comments
	}
	return err, comments
}

func GiveComment(comment model.CommentArea) error {
	err := dao.GiveComment(comment)
	if err != nil {
		return err
	}
	return err
}
