package service

import (
	"database/sql"
	"douban/dao"
	"douban/model"
)

func UploadAvatar(user model.UserMenu) error {
	err := dao.UploadAvatar(user)
	if err != nil {
		return err
	}
	return err
}

func SetQuestion(user model.UserEncrypted) (error, bool) {
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

func UserSeen(user model.UserSeen) error {
	err := dao.UserSeen(user)
	if err != nil {
		return err
	}
	return err
}

func UserWantSee(user model.UserWantSee) error {
	err := dao.UserWantSee(user)
	if err != nil {
		return err
	}
	return err
}

func GetUserComment(username string) (error, []model.ShortReview, []model.MovieReview) {
	return dao.GetUserComment(username)
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

func CommentMovie(movie model.MovieReview) error {
	return dao.CommentMovie(movie)
}

func Comment(movie model.ShortReview) error {
	return dao.Comment(movie)
}

func SetIntroduce(user model.UserMenu) error {
	return dao.SetIntroduce(user)
}

func GetUserMenu(username string) (error, model.UserMenu) {
	return dao.UserMenuInfo(username)
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

func CheckPassword(user model.User) (error, bool) {
	err, check := dao.CheckPassword(user)
	if err != nil {
		return err, false
	}
	err, res := Interpretation(check, user.Password)
	return err, res
}

func CheckUsername(user model.User) (error, bool) {
	err := dao.CheckUsername(user)
	if err != nil {
		if err == sql.ErrNoRows {
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
