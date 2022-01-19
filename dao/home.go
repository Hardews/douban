package dao

import "douban/modle"

func Find(keyWord string) (error, []modle.MovieInfo) {
	var movies []modle.MovieInfo
	sqlStr := "select * from movieBaseInfo where ChineseName like ?"
	keyWord = "%" + keyWord + "%"

	rows, err := dB.Query(sqlStr, keyWord)
	if err != nil {
		return err, movies
	}
	defer rows.Close()

	for rows.Next() {
		var movie modle.MovieInfo
		err := rows.Scan(&movie.Num, &movie.Name, &movie.OtherName, &movie.Score, &movie.Area, &movie.Year, &movie.Starring, &movie.Director, &movie.Types)
		if err != nil {
			return err, movies
		}
		movies = append(movies, movie)
	}
	return err, movies
}
