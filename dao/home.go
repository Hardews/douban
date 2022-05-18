package dao

import "douban/model"

func Find(keyWord string) (error, []model.MovieInfo) {
	var movies []model.MovieInfo

	tx := dB.Where("name LIKE ?", "%"+keyWord+"%").Find(&[]model.MovieInfo{}).Scan(&movies)
	if tx.Error != nil {
		return tx.Error, nil
	}

	return nil, movies
}
