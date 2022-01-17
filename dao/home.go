package dao

import "douban/modle"

func GetMovieBaseInfo() (error, []modle.MovieBaseInfo) {
	var info modle.MovieBaseInfo
	var infos []modle.MovieBaseInfo
	sqlStr := "select * from movieBaseInfo"
	rows, err := dB.Query(sqlStr)
	if err != nil {
		return err, infos
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(sqlStr, &info)
		if err != nil {
			return err, infos
		}
		infos = append(infos, info)
	}
	return err, infos
}
