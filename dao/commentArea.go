package dao

import "douban/modle"

func GiveCommentLike(username string, movieNum, areaNum, num int) {

}

func GiveTopicLike(username string, movieNum, num int) error {
	sqlStr := "insert commentLike (username,movieNum,topicNum) values (?,?,?)"
	_, err := dB.Exec(sqlStr, username, movieNum, num)
	if err != nil {
		return err
	}

	var likeNum int
	sqlStr = "select likeNum from commentArea where movieNum = ? and num = ?"
	err = dB.QueryRow(sqlStr, movieNum, num).Scan(&likeNum)
	if err != nil {
		return err
	}
	likeNum = likeNum + 1

	sqlStr = "update commentArea set commentLike = ?"
	_, err = dB.Exec(sqlStr, likeNum)
	if err != nil {
		return err
	}
	return err
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
	sqlStr2 := "select username,comment,time,likeNum from comment where movieNum = ? and num = ?"
	rows, err := dB.Query(sqlStr2, movieNum, areaNum)
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
	sqlStr1 := "select num,username,topic,time,likeNum,commentNum from commentArea where movieNum = ?"
	rows, err := dB.Query(sqlStr1, movieNum)
	if err != nil {
		return err, commentTopics
	}
	defer rows.Close()

	for rows.Next() {
		var commentTopic modle.CommentArea
		err = rows.Scan(&commentTopic.Num, &commentTopic.Username, &commentTopic.Topic, &commentTopic.Time, &commentTopic.CommentNum)
		if err != nil {
			return err, commentTopics
		}
		commentTopics = append(commentTopics, commentTopic)
	}
	return err, commentTopics
}
