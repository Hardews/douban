package dao

import (
	"douban/model"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func UploadAvatar(user model.UserMenu) error {
	tx := dB.Model(&model.UserMenu{}).Where("username = ?", user.Username).Update("avatar", user.Avatar)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func SetQuestion(user model.UserEncrypted) (error, bool) {
	// 判断是否有该用户
	_, err := SelectQuestion(user.Username)
	switch {
	case err != nil && err == gorm.ErrRecordNotFound:
		err = nil
	case err == nil:
		return err, false
	}

	// 创建密保
	t := dB.Create(&user)
	if err := t.Error; err != nil {
		return err, false
	}
	return nil, true
}

func CheckAnswer(username string) (error, string) {
	var question model.UserEncrypted
	tx := dB.Where("username = ?", username).First(&model.UserEncrypted{}).Scan(&question)
	if err := tx.Error; err != nil {
		return err, question.Answer
	}
	return nil, question.Answer
}

func SelectQuestion(username string) (string, error) {
	var user model.UserEncrypted
	tx := dB.Where("username = ?", username).First(&model.UserEncrypted{}).Scan(&user)
	if tx.Error != nil {
		return "", tx.Error
	}
	return user.Question, nil
}

func ChangePassword(user model.User) error {
	tx := dB.Model(&model.User{}).Where("username = ?", user.Username).Update("password", user.Password)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func CheckPassword(user model.User) (error, string) {
	var check model.User
	tx := dB.Where("username = ?", user.Username).First(&model.User{}).Scan(&check)
	if err := tx.Error; err != nil {
		return err, check.Password
	}
	return nil, check.Password
}

func CheckUsername(user model.User) error {
	var uUser model.User
	tx := dB.Where("username = ?", user.Username).First(&model.User{}).Scan(&uUser)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func WriteIn(user model.User) error {
	tx := dB.Begin()

	t := tx.Create(&user)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err
	}

	// 创建个人信息初始数据
	var userMenu model.UserMenu
	userMenu.Username = user.Username
	t = tx.Create(&userMenu)
	if err := t.Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
