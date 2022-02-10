package service

import (
	"database/sql"
	"douban/dao"
	"douban/modle"
)

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
		err := dao.UpdateComment(username, txt, movieNum, areaNum)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUserWantSee(username string) (error, []modle.UserHistory) {
	err, wantSee := dao.GetWantSee(username)
	if err != nil {
		return err, wantSee
	}
	return err, wantSee
}

func GetUserSeen(username string) (error, []modle.UserHistory) {
	err, seen := dao.GetSeen(username)
	if err != nil {
		return err, seen
	}
	return err, seen
}

func UserSeen(username, comment, label string, movieNum int) error {
	err := dao.UserSeen(username, comment, label, movieNum)
	if err != nil {
		return err
	}
	return err
}

func UserWantSee(username, comment, label string, movieNum int) error {
	err := dao.UserWantSee(username, comment, label, movieNum)
	if err != nil {
		return err
	}
	return err
}

func GetUserComment(username string) (error, []modle.UserComment, []modle.UserComment) {
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

func CommentMovie(Txt, username, commentTopic string, movieNum int) error {
	err := dao.CommentMovie(Txt, username, commentTopic, movieNum)
	if err != nil {
		return err
	}
	return err
}

func Comment(Txt, username string, movieNum int) error {
	err := dao.Comment(Txt, username, movieNum)
	if err != nil {
		return err
	}
	return err
}

func SetIntroduce(username, introduce string) error {
	err := dao.SetIntroduce(username, introduce)
	if err != nil {
		return err
	}
	return err
}

func GetUserMenu(username string) (error, modle.User) {
	err, user := dao.UserMenuInfo(username)
	if err != nil {
		return err, user
	}
	return err, user
}

func ChangePassword(user modle.User) error {
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

func CheckPassword(user modle.User) (error, bool) {
	err, check := dao.CheckPassword(user)
	if err != nil {
		return err, false
	}
	err, res := Interpretation(check.Password, user.Password)
	return err, res
}

func CheckUsername(user modle.User) (error, bool) {
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

func WriteIn(user modle.User) error {
	err := dao.WriteIn(user)
	if err != nil {
		return err
	}
	return err
}
