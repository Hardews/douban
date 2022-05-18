package dao

import (
	"douban/model"
)

func GetWantSee(username string) (error, []model.UserWantSee) {
	var wantSees []model.UserWantSee
	tx := dB.Where("username = ?", username).Find(&[]model.UserWantSee{}).Scan(&wantSees)
	if tx.Error != nil {
		return tx.Error, wantSees
	}
	return nil, wantSees
}

func GetSeen(username string) (error, []model.UserSeen) {
	var Seen []model.UserSeen
	TX := dB.Where("username = ?", username).Find(&[]model.UserSeen{}).Scan(&Seen)
	if TX.Error != nil {
		return TX.Error, Seen
	}
	return nil, Seen
}

func UserSeen(seen model.UserSeen) error {
	// 开启事务
	tx := dB.Begin()

	// 创建看过记录
	tx.Create(&seen)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	// 获取该电影的看过人数
	var movieExtra model.MovieExtra
	t := tx.Where("num = ?", seen.Num).Find(&model.MovieExtra{}).Scan(&movieExtra)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err
	}

	// 人数加一并更新
	movieExtra.SeenNum += 1
	t = tx.Model(&movieExtra).Update("seen_num", movieExtra.SeenNum)
	if t.Error != nil {
		tx.Rollback()
		return t.Error
	}

	// 无错误提交
	tx.Commit()
	return nil
}

func UserWantSee(want model.UserWantSee) error {
	// 开启事务
	tx := dB.Begin()

	// 创建想看记录
	tx.Create(&want)
	if err := tx.Error; err != nil {
		// 出错回滚事务
		tx.Rollback()
		return err
	}

	// 获取该电影的想看人数
	var movieExtra model.MovieExtra
	t := tx.Where("num = ?", want.Num).Find(&model.MovieExtra{}).Scan(&movieExtra)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err
	}

	// 人数加一并更新
	movieExtra.WantSeeNum += 1
	t = tx.Model(&movieExtra).Update("want_see_num", movieExtra.WantSeeNum)
	if t.Error != nil {
		tx.Rollback()
		return t.Error
	}

	// 无错误提交
	tx.Commit()
	return nil
}

func GetUserComment(username string) (error, []model.ShortReview, []model.MovieReview) {
	var short []model.ShortReview
	var long []model.MovieReview
	tx := dB.Where("username = ?", username).Find(&[]model.ShortReview{}).Scan(&short)
	if tx.Error != nil {
		return tx.Error, short, long
	}
	tx = dB.Where("username = ?", username).Find(&[]model.MovieReview{}).Scan(&long)
	if tx.Error != nil {
		return tx.Error, short, long
	}
	return nil, short, long
}

func SetIntroduce(user model.UserMenu) error {
	return dB.Model(&model.UserMenu{}).Where("username = ?", user.Username).Update("introduce", user.Introduce).Error
}

func UserMenuInfo(username string) (error, model.UserMenu) {
	var user model.UserMenu
	tx := dB.Where("username = ?", username).First(&model.UserMenu{}).Scan(&user)
	if err := tx.Error; err != nil {
		return err, user
	}
	return nil, user
}
