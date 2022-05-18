package main

import (
	"douban/dao"
	"douban/model"
	"fmt"
	"github.com/MashiroC/begonia"
	"github.com/MashiroC/begonia/app/option"
	"time"
)

func main() {
	// 一般情况下，addr是服务中心的地址。
	s := begonia.NewServer(option.Addr(":12306"))

	// 会通过反射的方式把service结构体下面所有公开的方法注册到LoginServer服务上。
	s.Register("AdmServer", &service{})

	// 让服务器持续睡眠，不然service会因为主进程退出而直接结束。
	for {
		time.Sleep(1 * time.Hour)
	}
}

type service struct{}

// NewMovie 函数的参数和返回值会被反射解析，注册为一个远程函数。
// 注册的函数没有特定的格式和写法。
func (*service) NewMovie(movie model.MovieInfo) (bool, int) {
	err, num := dao.NewMovie(movie)
	if err != nil {
		fmt.Println("set new movie failed,err:", err)
		return false, 0
	}
	return true, num
}
