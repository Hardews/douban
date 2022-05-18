package service

import (
	"douban/dao"
	"douban/model"
	"gorm.io/gorm"
)

func DoNotLikeTopic(user model.TopicLike) error {
	return dao.DoNotLikeTopic(user)
}

func DoNotLikeComment(user model.CommentLike) error {
	return dao.DoNotLikeComment(user)
}

func DeleteComment(user model.Comment) (error, bool) {
	err := dao.DeleteComment(user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return err, false
		}
		return err, false
	}
	return err, true
}

func DeleteCommentArea(userOp model.CommentArea) (error, bool) {
	err := dao.DeleteCommentArea(userOp)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return err, false
		}
		return err, false
	}
	return err, true
}

func GiveCommentLike(likeUser model.CommentLike) (error, bool) {
	return dao.GiveCommentLike(likeUser)
}

func GiveTopicLike(userLike model.TopicLike) (error, bool) {
	return dao.GiveTopicLike(userLike)
}

func GetCommentArea(movieNum int) (error, []model.CommentArea) {
	return dao.GetCommentArea(movieNum)
}

func GetCommentByNum(areaNum uint) (error, []model.Comment) {
	return dao.GetCommentByNum(areaNum)
}

func GiveComment(comment model.Comment) error {
	return dao.GiveComment(comment)
}
