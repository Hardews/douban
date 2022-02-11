package dao

func Find(keyWord string) (error, []int) {
	var movies []int
	sqlStr := "select Num from movie_Base_Info where ChineseName like ? or otherName like ? or types like ? "
	keyWord = "%" + keyWord + "%"

	rows, err := dB.Query(sqlStr, keyWord, keyWord, keyWord)
	if err != nil {
		return err, movies
	}
	defer rows.Close()

	for rows.Next() {
		var movie int
		err := rows.Scan(&movie)
		if err != nil {
			return err, movies
		}
		movies = append(movies, movie)
	}
	return err, movies
}
