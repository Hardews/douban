package dao

import (
	"douban/model"
)

func NewMovie(movie model.MovieInfo) (error, int) {
	tx := dB.Create(&movie)
	if err := tx.Error; err != nil {
		return err, 0
	}

	var movies model.MovieInfo
	dB.Last(&model.MovieInfo{}).Scan(&movies)

	return nil, movies.Num
}
