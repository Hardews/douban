package dao

import (
	"douban/model"
)

func DeleteShortComment(username string, movieNum int) error {
	tx := dB.Begin()

	t := tx.Where("username = ? and movie_num = ?", username, movieNum).Delete(&model.ShortReview{})
	if t.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	tx.Commit()

	return nil
}

func DeleteLongComment(username string, movieNum int) error {
	tx := dB.Begin()

	t := tx.Where("username = ? and movie_num = ?", username, movieNum).Delete(&model.MovieReview{})
	if t.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	var num int
	t = tx.Select("comment_num").Where("num = ?", movieNum).First(&num).Scan(&num)
	if t.Error != nil {
		tx.Rollback()
		return t.Error
	}

	num -= 1
	t = tx.Select("comment_num").Where("num = ?", movieNum).Create(&num)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func DeleteSeen(seen model.UserSeen) error {
	tx := dB.Begin()
	t := tx.Where("username = ? AND movie_num = ?", seen.Username, seen.Num).Delete(&seen)
	if t.Error != nil {
		tx.Rollback()
		return t.Error
	}

	// 查找想看数
	var seenNum int
	tx.Select("seen_num").Where("num = ?", seen.Num).First(&seenNum).Scan(&seenNum)
	if err := tx.Select("seen_num").Where("num = ?", seen.Num).First(&seenNum).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 减一
	seenNum -= 1
	t = tx.Select("seen_num").Create(&seenNum)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func DeleteWantSee(want model.UserWantSee) error {
	tx := dB.Begin()
	t := tx.Where("username = ? AND movie_num = ?", want.Username, want.Num).Delete(&want)
	if t.Error != nil {
		tx.Rollback()
		return t.Error
	}

	// 查找想看数
	var wantSeeNum int
	t = tx.Select("want_see_num").Where("num = ?", want.Num).First(&wantSeeNum).Scan(&wantSeeNum)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err
	}

	// 减一
	wantSeeNum -= 1
	t = tx.Select("want_see_num").Create(&wantSeeNum)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func GetComment(num int) (error, []model.MovieReview) {
	var comments []model.MovieReview
	tx := dB.Where("movie_num = ?", num).Find(&model.MovieReview{}).Scan(&comments)
	if tx.Error != nil {
		return tx.Error, comments
	}
	return nil, comments
}

func GetMovieComment(num int) (error, []model.ShortReview) {
	var comments []model.ShortReview
	tx := dB.Where("movie_num = ?", num).Find(&[]model.ShortReview{}).Scan(&comments)
	if tx.Error != nil {
		return tx.Error, comments
	}
	return nil, comments
}

func UpdateShortComment(username, txt string, movieNum int) error {
	tx := dB.Model(&model.ShortReview{}).Select("txt").Where("username = ? AND movie_num = ?", username, movieNum).Update("txt", txt)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func SelectShortComment(username string, movieNum int) (error, bool) {
	var review model.ShortReview
	tx := dB.Where("username = ? AND movie_num = ?", username, movieNum).First(&model.ShortReview{}).Scan(&review)
	if tx.Error != nil {
		return tx.Error, false
	}
	return nil, true
}

func Comment(shortReview model.ShortReview) error {
	tx := dB.Create(&shortReview)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func UpdateLongComment(username, txt string, movieNum int) error {
	tx := dB.Model(&model.MovieReview{}).Select("txt").Where("username = ? AND movie_num = ?", username, movieNum).Update("txt", txt)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func SelectLongComment(username string, movieNum int) (error, bool) {
	var review model.MovieReview
	tx := dB.Where("username = ? AND movie_num = ?", username, movieNum).First(&model.MovieReview{}).Scan(&review)
	if tx.Error != nil {
		return tx.Error, false
	}
	return nil, true
}

func CommentMovie(movieReview model.MovieReview) error {
	t := dB.Create(&movieReview)
	if err := t.Error; err != nil {
		return err
	}
	return nil
}

func FindWithCategory(category string) (error, []model.MovieInfo) {
	var movies []model.MovieInfo
	tx := dB.Where("types = ?", category).Find(&[]model.MovieInfo{}).Scan(&movies)
	if err := tx.Error; err != nil {
		return err, movies
	}
	return nil, movies
}

func GetAMovieInfo(movieNum int) (error, model.MovieInfo) {
	var movie model.MovieInfo
	tx := dB.Where("num = ?", movieNum).First(&model.MovieInfo{}).Scan(&movie)
	if tx.Error != nil {
		return tx.Error, movie
	}
	return nil, movie
}
