package dao

import "douban/modle"

func GetMovieBaseInfo() (error, []modle.MovieInfo) {
	var info modle.MovieInfo
	var infos []modle.MovieInfo
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

func GetMovieAllInfo() (error, []modle.MovieInfo) {
	var infos []modle.MovieInfo

	sqlStr := "select * from movieBaseInfo and movieInfo"
	rows, err := dB.Query(sqlStr)
	if err != nil {
		return err, infos
	}

	defer rows.Close()

	for rows.Next() {
		var info modle.MovieInfo

		err = rows.Scan(sqlStr, &info)
		if err != nil {
			return err, infos
		}

		infos = append(infos, info)
	}
	return err, infos
}
