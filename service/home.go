package service

import (
	"douban/dao"
)

func Find(keyword string) (error, []int) {
	err, movieNums := dao.Find(keyword)
	if err != nil {
		return err, movieNums
	}
	return err, movieNums
}
