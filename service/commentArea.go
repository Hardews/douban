package service

import (
	"douban/dao"
	"douban/modle"
)

func GiveTopicLike(username string, movieNum, num int) error {
	err := dao.GiveTopicLike(username, movieNum, num)
	if err != nil {
		return err
	}
	return err
}

func GetCommentArea(movieNum int) (error, []modle.CommentArea) {
	err, commentAreas := dao.GetCommentArea(movieNum)
	if err != nil {
		return err, commentAreas
	}
	return err, commentAreas
}

func GetCommentByNum(movieNum, areaNum int) (error, []modle.CommentArea) {
	err, comments := dao.GetCommentByNum(movieNum, areaNum)
	if err != nil {
		return err, comments
	}
	return err, comments
}

func GiveComment(comment modle.CommentArea) error {
	err := dao.GiveComment(comment)
	if err != nil {
		return err
	}
	return err
}
