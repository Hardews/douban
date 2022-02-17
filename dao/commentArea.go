package dao

import (
	"database/sql"
	"douban/modle"
)

func SelectArea(username string, movieNum int) (error, bool, int) {
	var num int
	sqlStr := "select num from comment_Area where username = ? and movieNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, false, 0
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, movieNum).Scan(&num)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, true, 0
		}
		return err, false, 0
	}
	return err, true, num
}

func SelectComment(username string, movieNum, areaNum int) (error, bool, int) {
	var num int
	sqlStr := "select no from comment where username = ? and movieNum = ? and areaNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, false, 0
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, movieNum, areaNum).Scan(&num)
	if err != nil {
		if err == sql.ErrNoRows {
			return err, true, 0
		}
		return err, false, 0
	}
	return err, true, num
}

func UpdateComment(username, txt string, movieNum, areaNum int) error {
	sqlStr := "update comment set txt=? where movieNum = ? and username = ? and areaNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(txt, movieNum, username, areaNum)
	if err != nil {
		return err
	}
	return err
}

func UpdateCommentArea(username, txt string, movieNum int) error {
	sqlStr := "update comment_Area set topic=? where movieNum = ? and username = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(txt, movieNum, username)
	if err != nil {
		return err
	}

	err, _, areaNum := SelectArea(username, movieNum)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err
		}
		return err
	}

	sqlStr = "DELETE FROM comment where areaNum = ? and movieNum = ?;"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(areaNum, movieNum)
	if err != nil {
		return err
	}
	return err
}

func DoNotLikeTopic(username string, areaNum int) error {
	var likeNum int
	sqlStr := "select likeNum from comment_Area where num = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(areaNum).Scan(&likeNum)
	if err != nil {
		return err
	}

	var iUsername string
	sqlStr = "select username from topic_Like where username = ? and topicNum = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(username, areaNum).Scan(&iUsername)
	if err != nil {
		return err
	}

	sqlStr = "delete topic_Like where username = ? and topicNum = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username, areaNum)
	if err != nil {
		return err
	}

	likeNum = likeNum - 1
	sqlStr = "update comment_Area set likeNum = ? where num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(likeNum, areaNum)
	return err
}

func DoNotLikeComment(username string, areaNum, commentNum int) error {
	var likeNum int
	sqlStr := "select likeNum from comment where num = ? and no = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(areaNum, commentNum).Scan(&likeNum)
	if err != nil {
		return err
	}

	var iUsername string
	sqlStr = "select username from comment_Like where username = ? and topicNum = ? and commentNum = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(username, areaNum, commentNum).Scan(&iUsername)
	if err != nil {
		return err
	}

	sqlStr = "delete comment_Like where username = ? and topicNum = ? and commentNum = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username, areaNum, commentNum)
	if err != nil {
		return err
	}

	likeNum = likeNum - 1
	sqlStr = "update comment set likeNum = ? where num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(likeNum, areaNum)
	return err
}

func DeleteComment(username string, movieNum, areaNum int) error {
	var iMovieNum, iAreaNum, iCommentNum string
	sqlStr := "select movieNum,num,no from comment where movieNum = ? and areaNum = ? and username = ? "
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(movieNum, areaNum, username).Scan(&iMovieNum, &iAreaNum, &iCommentNum)
	if err != nil {
		return err
	}

	iMovieNum, iAreaNum, iCommentNum = iMovieNum+"已删除", iAreaNum+"已删除", iCommentNum+"已删除"
	sqlStr = "update comment set movieNum = ?,num = ?,no = ? where movieNum = ? and areaNum = ? and username = ? "
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(iMovieNum, iAreaNum, iCommentNum, movieNum, areaNum, username)
	if err != nil {
		return err
	}
	return err
}

func DeleteCommentArea(movieNum, areaNum int) error {
	var iAreaNum string
	sqlStr := "select username from comment_Area where movieNum = ? and areaNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(movieNum, areaNum).Scan(&iAreaNum)
	if err != nil {
		return err
	}

	iAreaNum = iAreaNum + "已删除"
	sqlStr = "update comment_Area set areaNum = ? where movieNum = ? and areaNum = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(iAreaNum, movieNum, areaNum)
	if err != nil {
		return err
	}
	return err
}

func GiveCommentLike(username, name string, movieNum, areaNum int) (error, bool) {
	var iUsername string
	sqlStr := "select username from comment_Like where  movieNum = ? and topicNum = ? and username = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, false
	}
	err = stmt.QueryRow(movieNum, areaNum, username).Scan(&iUsername)
	switch err {
	case nil:
		return err, false
	case sql.ErrNoRows:
		err = nil
	default:
		return err, false
	}

	sqlStr = "insert comment_Like (username,movieNum,topicNum) values (?,?,?)"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err, false
	}

	_, err = stmt.Exec(username, movieNum, areaNum)
	if err != nil {
		return err, false
	}

	var likeNum int
	sqlStr = "select likeNum from comment where movieNum = ? and num = ? and username = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err, false
	}

	err = stmt.QueryRow(movieNum, areaNum, name).Scan(&likeNum)
	if err != nil {
		return err, false
	}
	likeNum = likeNum + 1

	sqlStr = "update comment set likeNum = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err, false
	}
	defer stmt.Close()

	_, err = stmt.Exec(likeNum)
	if err != nil {
		return err, false
	}
	return err, true
}

func GiveTopicLike(username string, movieNum, num int) (error, bool) {
	var iUsername string
	sqlStr := "select username from topic_Like where username = ? and movieNum = ? and topicNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, false
	}

	err = stmt.QueryRow(username, movieNum, num).Scan(&iUsername)
	switch err {
	case nil:
		return err, false
	case sql.ErrNoRows:
		err = nil
	default:
		return err, false
	}

	sqlStr = "insert topic_Like (username,movieNum,topicNum) values (?,?,?)"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err, false
	}

	_, err = stmt.Exec(username, movieNum, num)
	if err != nil {
		return err, false
	}

	var likeNum int
	sqlStr = "select likeNum from comment_Area where movieNum = ? and num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err, false
	}

	err = stmt.QueryRow(movieNum, num).Scan(&likeNum)
	if err != nil {
		return err, false
	}
	likeNum = likeNum + 1

	sqlStr = "update comment_Area set LikeNum = ? where movieNum = ? and num = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err, false
	}
	defer stmt.Close()

	_, err = stmt.Exec(likeNum, movieNum, num)
	if err != nil {
		return err, false
	}
	return err, true
}

func GiveComment(comment modle.CommentArea) error {
	sqlStr := "insert comment (areaNum,username,txt,movieNum) values (?,?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(comment.Num, comment.Username, comment.Comment, comment.MovieNum)
	if err != nil {
		return err
	}

	var commentNum int
	sqlStr = "select comment_Num from comment_Area where num = ? and movieNum = ?"
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(comment.Num, comment.MovieNum).Scan(&commentNum)
	if err != nil {
		return err
	}

	commentNum = commentNum + 1
	sqlStr = "update comment_Area set commentNum = ? where num = ? and movieNum = ? "
	stmt, err = dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(commentNum, comment.Num, comment.MovieNum)
	if err != nil {
		return err
	}

	return err
}

func SetCommentArea(username, topic string, movieNum int) error {
	sqlStr := "insert comment_Area (username,topic,movieNum) values (?,?,?)"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, topic, movieNum)
	if err != nil {
		return err
	}
	return err
}

func GetCommentByNum(movieNum, areaNum int) (error, []modle.CommentArea) {
	var comments []modle.CommentArea
	sqlStr := "select username,txt,time,likeNum from comment where movieNum = ? and areaNum = ?"
	stmt, err := dB.Prepare(sqlStr)
	if err != nil {
		return err, comments
	}
	defer stmt.Close()

	rows, err := stmt.Query(movieNum, areaNum)
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

func GetCommentArea(movieNum int) (error, []modle.CommentArea) {
	var commentTopics []modle.CommentArea
	sqlStr1 := "select num,username,topic,time,likeNum,commentNum from comment_Area where movieNum = ?"
	stmt, err := dB.Prepare(sqlStr1)
	if err != nil {
		return err, commentTopics
	}
	defer stmt.Close()

	rows, err := stmt.Query(movieNum)
	if err != nil {
		return err, commentTopics
	}
	defer rows.Close()

	for rows.Next() {
		var commentTopic modle.CommentArea
		err := rows.Scan(&commentTopic.Num, &commentTopic.Username, &commentTopic.Topic, &commentTopic.Time, &commentTopic.LikeNum, &commentTopic.CommentNum)
		if err != nil {
			return err, commentTopics
		}

		commentTopics = append(commentTopics, commentTopic)
	}

	return err, commentTopics
}
