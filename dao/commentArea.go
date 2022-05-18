package dao

import (
	"douban/model"
	"gorm.io/gorm"
)

func SelectArea(username string, movieNum int) (error, bool, int) {
	var num int
	tx := dB.Model(&model.CommentArea{}).Select("num").Where("username = ? AND movie_num = ?", username, movieNum).Scan(&num)
	if err := tx.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return err, true, 0
		}
		return err, false, 0
	}
	return nil, true, num
}

func SelectComment(username string, movieNum, areaNum int) (error, bool, int) {
	var num int
	tx := dB.Model(&model.Comment{}).Select("num").Where("username = ? AND movie_num = ? AND comment_id = ?", username, movieNum, areaNum).Scan(&num)
	if err := tx.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err, true, 0
		}
		return err, false, 0
	}
	return nil, true, num
}

func UpdateComment(username, txt string, areaNum int) error {
	tx := dB.Model(&model.Comment{}).Where("username = ? AND comment_id = ?", username, areaNum).Update("txt", txt)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func UpdateCommentArea(username, txt string, movieNum int) error {
	tx := dB.Model(&model.CommentArea{}).Where("username = ? AND movie_num = ?", username, movieNum).Update("txt", txt)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func DoNotLikeTopic(likeUser model.TopicLike) error {
	// 事务开启
	tx := dB.Begin()
	var user model.TopicLike

	// 检查是否点过赞
	t := tx.Where("username = ? AND topic_id = ?", likeUser.Username, likeUser.TopicId).Find(&model.TopicLike{}).Scan(&user)
	if t.Error != nil {
		if t.Error == gorm.ErrRecordNotFound {
			t.Error = nil
		}
		return t.Error
	}

	// 删除点赞记录
	t = tx.Delete(&likeUser)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err
	}

	// 对点赞数进行减一并保存
	var likeNum int
	t = tx.Model(&model.CommentArea{}).Select("like_num").Where("id = ? AND topic_id = ?", likeUser.Username, likeUser.TopicId).Scan(&likeNum)
	if t.Error != nil {
		tx.Rollback()
		return t.Error
	}

	likeNum -= 1
	t = tx.Model(&model.CommentArea{}).Select("like_num").Where("id = ? AND topic_id = ?", likeUser.Username, likeUser.TopicId).Update("like_num", likeNum)
	if t.Error != nil {
		tx.Rollback()
		return t.Error
	}
	// 无误后提交
	tx.Commit()
	return nil
}

func DoNotLikeComment(likeUser model.CommentLike) error {
	// 事务开启
	tx := dB.Begin()
	var user model.CommentLike

	// 检查是否点过赞
	t := tx.Where("username = ? AND comment_id", likeUser.Username, likeUser.CommentId).Find(&model.CommentLike{}).Scan(&user)
	if t.Error != nil {
		if t.Error == gorm.ErrRecordNotFound {
			t.Error = nil
			return nil
		}
		return t.Error
	}

	// 删除点赞记录
	t = tx.Where("username = ? AND comment_id", likeUser.Username, likeUser.CommentId).Delete(&likeUser)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err
	}

	// 对点赞数进行减一并保存
	var likeNum int
	t = tx.Model(&model.Comment{}).Select("username = ? AND comment_id", likeUser.Username, likeUser.CommentId).First(&likeNum).Scan(&likeNum)
	if t.Error != nil {
		tx.Rollback()
		return t.Error
	}

	likeNum -= 1
	t = tx.Model(&model.Comment{}).Select("username = ? AND comment_id", likeUser.Username, likeUser.CommentId).Update("like_num", likeNum)
	if t.Error != nil {
		tx.Rollback()
		return t.Error
	}
	// 无误后提交
	tx.Commit()
	return nil
}

func DeleteComment(user model.Comment) error {
	tx := dB.Delete(&user)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func DeleteCommentArea(userOp model.CommentArea) error {
	tx := dB.Delete(&userOp)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func GiveCommentLike(likeUser model.CommentLike) (error, bool) {
	// 事务开启
	tx := dB.Begin()
	var user model.CommentLike

	// 检查是否点过赞
	t := tx.Where("username = ? AND movie_num = ? AND comment_id", likeUser.Username, likeUser.MovieNum, likeUser.CommentId).Find(&model.CommentLike{}).Scan(&user)
	if t.Error != nil {
		if t.Error == gorm.ErrRecordNotFound {
			t.Error = nil
		}
		return t.Error, false
	}

	// 如果有记录则是点过赞的，返回false
	if user.Username != "" {
		return nil, false
	}

	// 创建点赞记录
	t = tx.Create(&likeUser)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err, false
	}

	// 对点赞数进行加一并保存
	var likeNum int
	t = tx.Model(&model.Comment{}).Select("like_num").Where("id = ? AND movie_num = ? AND comment_id", likeUser.Username, likeUser.MovieNum, likeUser.CommentId).First(&likeNum).Scan(&likeNum)
	if t.Error != nil {
		tx.Rollback()
		return t.Error, false
	}

	likeNum += 1
	t = tx.Model(&model.Comment{}).Select("like_num").Where("id = ? AND movie_num = ? AND comment_id", likeUser.Username, likeUser.MovieNum, likeUser.CommentId).Update("like_num", likeNum)
	if t.Error != nil {
		tx.Rollback()
		return t.Error, false
	}
	// 无误后提交
	tx.Commit()
	return nil, true
}

func GiveTopicLike(likeUser model.TopicLike) (error, bool) {
	// 事务开启
	tx := dB.Begin()
	var user model.TopicLike

	// 检查是否点过赞
	t := tx.Where("username = ? AND movie_num = ?", likeUser.Username, likeUser.MovieNum).Scan(&user)
	if t.Error != nil {
		if t.Error == gorm.ErrRecordNotFound {
			t.Error = nil
		}
		return t.Error, false
	}

	// 如果有记录则是点过赞的，返回false
	if user.Username != "" {
		return nil, false
	}

	// 创建点赞记录
	t = tx.Create(&likeUser)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err, false
	}

	// 对点赞数进行加一并保存
	var likeNum int
	t = tx.Model(&model.CommentArea{}).Select("like_num").Where("id = ? AND movie_num = ?", likeUser.TopicId, likeUser.MovieNum).First(&likeNum).Scan(&likeNum)
	if t.Error != nil {
		tx.Rollback()
		return t.Error, false
	}

	likeNum += 1
	t = tx.Model(&model.CommentArea{}).Select("like_num").Where("id = ? AND movie_num = ?", likeUser.TopicId, likeUser.MovieNum).Update("like_num", likeNum)
	if t.Error != nil {
		tx.Rollback()
		return t.Error, false
	}
	// 无误后提交
	tx.Commit()
	return nil, true
}

func GiveComment(comment model.Comment) error {
	tx := dB.Begin()
	t := tx.Create(&comment)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err
	}

	var commentNum int
	t = tx.Model(&model.CommentArea{}).Select("comment_num").Where("id = ?", comment.CommentId).Scan(&commentNum)
	if t.Error != nil {
		tx.Rollback()
		return t.Error
	}

	commentNum += 1
	t = tx.Model(&model.CommentArea{}).Where("id = ?", comment.CommentId).Update("comment_num", commentNum)
	if t.Error != nil {
		tx.Rollback()
		return t.Error
	}

	tx.Commit()
	return nil
}

func SetCommentArea(area model.CommentArea) error {
	tx := dB.Create(&area)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func GetCommentByNum(areaNum uint) (error, []model.Comment) {
	var comments []model.Comment
	t := dB.Where("comment_id = ?", areaNum).Find(&[]model.Comment{}).Scan(&comments)
	if err := t.Error; err != nil {
		return err, comments
	}
	return nil, comments
}

func GetCommentArea(movieNum int) (error, []model.CommentArea) {
	var commentTopics []model.CommentArea
	t := dB.Where("movie_num = ?", movieNum).Find(&[]model.CommentArea{}).Scan(&commentTopics)
	if err := t.Error; err != nil {
		return err, commentTopics
	}
	return nil, commentTopics
}
