package service

import (
	"douban/dao"
	"douban/model"
	"gorm.io/gorm"
)

func UploadAvatar(user model.UserMenu) error {
	err := dao.UploadAvatar(user)
	if err != nil {
		return err
	}
	return err
}

func SetQuestion(user model.UserEncrypted) (error, bool) {
	// 对答案进行加密
	var err error
	err, user.Answer = Encryption(user.Answer)
	if err != nil {
		return err, false
	}
	err, flag := dao.SetQuestion(user)
	if err != nil {
		return err, false
	}
	return err, flag
}

func CheckAnswer(username, answer string) (error, bool) {
	err, check := dao.CheckAnswer(username)
	if err != nil {
		return err, false
	}
	err, res := Interpretation(check, answer)
	return err, res
}

func SelectQuestion(username string) (string, error) {
	question, err := dao.SelectQuestion(username)
	if err != nil {
		return question, err
	}
	return question, err
}

func UpdateComment(username, txt string, movieNum, choose, areaNum int) error {
	switch choose {
	case 1:
		err := dao.UpdateLongComment(username, txt, movieNum)
		if err != nil {
			return err
		}
	case 2:
		err := dao.UpdateShortComment(username, txt, movieNum)
		if err != nil {
			return err
		}
	case 3:
		err := dao.UpdateCommentArea(username, txt, movieNum)
		if err != nil {
			return err
		}
	case 4:
		err := dao.UpdateComment(username, txt, areaNum)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUserWantSee(username string) (error, []model.UserWantSee) {
	err, wantSee := dao.GetWantSee(username)
	if err != nil {
		return err, wantSee
	}
	return err, wantSee
}

func GetUserSeen(username string) (error, []model.UserSeen) {
	err, seen := dao.GetSeen(username)
	if err != nil {
		return err, seen
	}
	return err, seen
}

func UserSeen(seen model.UserSeen) error {
	err := dao.UserSeen(seen)
	if err != nil {
		return err
	}
	return err
}

func UserWantSee(wantSee model.UserWantSee) error {
	err := dao.UserWantSee(wantSee)
	if err != nil {
		return err
	}
	return err
}

func GetUserComment(username string) (error, []model.ShortReview, []model.MovieReview) {
	err, shortComments, longComments := dao.GetUserComment(username)
	if err != nil {
		return err, shortComments, longComments
	}
	return err, shortComments, longComments
}

func SelectComment(username string, movieNum, choose, areaNum int) (error, bool, int) {
	switch choose {
	case 1:
		err, flag := dao.SelectLongComment(username, movieNum)
		if err != nil {
			return err, flag, 0
		}
		return err, flag, 0
	case 2:
		err, flag := dao.SelectShortComment(username, movieNum)
		if err != nil {
			return err, flag, 0
		}
		return err, flag, 0
	case 3:
		err, flag, num := dao.SelectArea(username, movieNum)
		if err != nil {
			return err, flag, 0
		}
		return err, flag, num
	case 4:
		err, flag, num := dao.SelectComment(username, movieNum, areaNum)
		if err != nil {
			return err, flag, 0
		}
		return err, flag, num
	}
	return nil, false, 0
}

func CommentMovie(movieReview model.MovieReview) error {
	err := dao.CommentMovie(movieReview)
	if err != nil {
		return err
	}
	return err
}

func Comment(shortReview model.ShortReview) error {
	err := dao.Comment(shortReview)
	if err != nil {
		return err
	}
	return err
}

func SetIntroduce(user model.UserMenu) error {
	err := dao.SetIntroduce(user)
	if err != nil {
		return err
	}
	return err
}

func GetUserMenu(username string) (error, model.UserMenu) {
	err, user := dao.UserMenuInfo(username)
	if err != nil {
		return err, user
	}
	return err, user
}

func ChangePassword(user model.User) error {
	var err error
	err, user.Password = Encryption(user.Password)
	if err != nil {
		return err
	}
	err = dao.ChangePassword(user)
	if err != nil {
		return err
	}
	return err
}

func CheckPassword(username, password string) (error, bool) {
	err, check := dao.CheckPassword(username)
	if err != nil {
		return err, false
	}
	err, res := Interpretation(check.Password, password)
	if err != nil {
		return err, false
	}
	return err, res
}

func CheckUsername(user model.User) (error, bool) {
	err := dao.CheckUsername(user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return err, true
		}
		return err, false
	}
	return err, false
}

func WriteIn(user model.User) error {
	err := dao.WriteIn(user)
	if err != nil {
		return err
	}
	return err
}
