package dao

import (
	"database/sql"
	"douban/modle"
)

func DoNotLikeTopic(username string, areaNum int) error {
	var likeNum int
	sqlStr := "select likeNum from commentArea where num = ?"
	err := dB.QueryRow(sqlStr, areaNum).Scan(&likeNum)
	if err != nil {
		return err
	}

	var iUsername string
	sqlStr = "select username from topicLike where username = ? and topicNum = ?"
	err = dB.QueryRow(sqlStr, username, areaNum).Scan(&iUsername)
	if err != nil {
		return err
	}

	sqlStr = "delete topicLike where username = ? and topicNum = ?"
	_, err = dB.Exec(sqlStr, username, areaNum)
	if err != nil {
		return err
	}

	likeNum = likeNum - 1
	sqlStr = "update commentArea set likeNum = ? where num = ?"
	_, err = dB.Exec(sqlStr, likeNum, areaNum)
	return err
}

func DoNotLikeComment(username string, areaNum, commentNum int) error {
	var likeNum int
	sqlStr := "select likeNum from comment where num = ? and no = ?"
	err := dB.QueryRow(sqlStr, areaNum, commentNum).Scan(&likeNum)
	if err != nil {
		return err
	}

	var iUsername string
	sqlStr = "select username from commentLike where username = ? and topicNum = ? and commentNum = ?"
	err = dB.QueryRow(sqlStr, username, areaNum, commentNum).Scan(&iUsername)
	if err != nil {
		return err
	}

	sqlStr = "delete commentLike where username = ? and topicNum = ? and commentNum = ?"
	_, err = dB.Exec(sqlStr, username, areaNum, commentNum)
	if err != nil {
		return err
	}

	likeNum = likeNum - 1
	sqlStr = "update comment set likeNum = ? where num = ?"
	_, err = dB.Exec(sqlStr, likeNum, areaNum)
	return err
}

func DeleteComment(movieNum, areaNum, commentNum int) error {
	var iMovieNum, iAreaNum, iCommentNum string
	sqlStr := "select movieNum,num,no from comment where movieNum = ? and num = ? and no = ?"
	err := dB.QueryRow(sqlStr, movieNum, areaNum, commentNum).Scan(&iMovieNum, &iAreaNum, &iCommentNum)
	if err != nil {
		return err
	}

	iMovieNum, iAreaNum, iCommentNum = iMovieNum+"已删除", iAreaNum+"已删除", iCommentNum+"已删除"
	sqlStr = "update comment set movieNum = ?,num = ?,no = ? where movieNum = ? and num = ? and no = ?"
	_, err = dB.Exec(sqlStr, iMovieNum, iAreaNum, iCommentNum, movieNum, areaNum, commentNum)
	if err != nil {
		return err
	}
	return err
}

func DeleteCommentArea(movieNum, areaNum int) error {
	var iAreaNum string
	sqlStr := "select username from commentArea where movieNum = ? and areaNum = ?"
	err := dB.QueryRow(sqlStr, movieNum, areaNum).Scan(&iAreaNum)
	if err != nil {
		return err
	}

	iAreaNum = iAreaNum + "已删除"
	sqlStr = "update commentArea set areaNum = ? where movieNum = ? and areaNum = ?"
	_, err = dB.Exec(sqlStr, iAreaNum, movieNum, areaNum)
	if err != nil {
		return err
	}
	return err
}

func GiveCommentLike(username string, movieNum, areaNum, commentNum int) (error, bool) {
	var iUsername string
	sqlStr := "select username from commentLike where commentNum = ? and movieNum = ? and topicNum = ? and username = ?"
	err := dB.QueryRow(sqlStr, commentNum, movieNum, areaNum, username).Scan(&iUsername)
	switch err {
	case nil:
		return err, false
	case sql.ErrNoRows:
		err = nil
	default:
		return err, false
	}

	sqlStr = "insert commentLike (username,movieNum,topicNum,commentNum) values (?,?,?,?)"
	_, err = dB.Exec(sqlStr, username, movieNum, areaNum, commentNum)
	if err != nil {
		return err, false
	}

	var likeNum int
	sqlStr = "select likeNum from comment where movieNum = ? and num = ? and no = ?"
	err = dB.QueryRow(sqlStr, movieNum, areaNum, commentNum).Scan(&likeNum)
	if err != nil {
		return err, false
	}
	likeNum = likeNum + 1

	sqlStr = "update comment set likeNum = ?"
	_, err = dB.Exec(sqlStr, likeNum)
	if err != nil {
		return err, false
	}
	return err, true
}

func GiveTopicLike(username string, movieNum, num int) (error, bool) {
	var iUsername string
	sqlStr := "select username from topicLike where username = ? and movieNum = ? and topicNum = ?"
	err := dB.QueryRow(sqlStr, username, movieNum, num).Scan(&iUsername)
	switch err {
	case nil:
		return err, false
	case sql.ErrNoRows:
		err = nil
	default:
		return err, false
	}

	sqlStr = "insert topicLike (username,movieNum,topicNum) values (?,?,?)"
	_, err = dB.Exec(sqlStr, username, movieNum, num)
	if err != nil {
		return err, false
	}

	var likeNum int
	sqlStr = "select likeNum from commentArea where movieNum = ? and num = ?"
	err = dB.QueryRow(sqlStr, movieNum, num).Scan(&likeNum)
	if err != nil {
		return err, false
	}
	likeNum = likeNum + 1

	sqlStr = "update commentArea set LikeNum = ? where movieNum = ? and num = ?"
	_, err = dB.Exec(sqlStr, likeNum, movieNum, num)
	if err != nil {
		return err, false
	}
	return err, true
}

func GiveComment(comment modle.CommentArea) error {
	sqlStr := "insert comment (num,username,txt,movieNum) values (?,?,?,?)"
	_, err := dB.Exec(sqlStr, comment.Num, comment.Username, comment.Comment, comment.MovieNum)
	if err != nil {
		return err
	}

	var commentNum int
	sqlStr = "select commentNum from commentArea where num = ? and movieNum = ?"
	err = dB.QueryRow(sqlStr, comment.Num, comment.MovieNum).Scan(&commentNum)
	if err != nil {
		return err
	}

	commentNum = commentNum + 1
	sqlStr = "update commentArea set commentNum = ? where num = ? and movieNum = ? "
	_, err = dB.Exec(sqlStr, commentNum, comment.Num, comment.MovieNum)
	if err != nil {
		return err
	}

	return err
}

func SetCommentArea(username, topic string, movieNum int) error {
	sqlStr := "insert commentArea (username,topic,movieNum) values (?,?,?)"
	_, err := dB.Exec(sqlStr, username, topic, movieNum)
	if err != nil {
		return err
	}
	return err
}

func GetCommentByNum(movieNum, areaNum int) (error, []modle.CommentArea) {
	var comments []modle.CommentArea
	sqlStr := "select username,txt,time,likeNum from comment where movieNum = ? and num = ?"
	rows, err := dB.Query(sqlStr, movieNum, areaNum)
	if err != nil {
		return err, comments
	}

	defer rows.Close()

	for rows.Next() {
		var comment modle.CommentArea
		err = rows.Scan(&comment.Username, &comment.Comment, &comment.Time, &comment.LikeNum)
		if err != nil {
			return err, comments
		}
		comments = append(comments, comment)
	}
	return err, comments
}

func GetCommentArea(movieNum, areaNum int) (error, modle.CommentArea) {
	var commentTopic modle.CommentArea
	sqlStr1 := "select num,username,topic,time,likeNum,commentNum from commentArea where movieNum = ? and num = ?"
	err := dB.QueryRow(sqlStr1, movieNum, areaNum).Scan(&commentTopic.Num, &commentTopic.Username, &commentTopic.Topic, &commentTopic.Time, &commentTopic.LikeNum, &commentTopic.CommentNum)
	if err != nil {
		return err, commentTopic
	}

	return err, commentTopic
}
