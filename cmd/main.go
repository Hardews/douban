package main

import (
	"douban/dao"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"net/http"
	"os"
	"strconv"
)

//写入文件
func SaveJokeFile(url, movieNum string) {

	//保存在本地的地址
	path := "D:/GOProjects/src/douban/movieFile/No " + movieNum + ".jpg"
	path1 := "/opt/gocode/src/douban/movieFile/No " + movieNum + ".jpg"
	err := dao.Write(path1, movieNum)
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}
	defer f.Close()

	//读取url的信息
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http err:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		//写入文件
		f.Write(buf[:n])
	}

}

func main() {
	dao.InitDB()
	fileName := "D://douban.xlsx"
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	for i, _ := range rows {
		if i == 251 {
			break
		}
		str := strconv.Itoa(i + 1)
		movieNum, err := f.GetCellValue("Sheet1", "A"+str)
		if err != nil {
			fmt.Println(err)
			return
		}
		img, err := f.GetCellValue("Sheet1", "B"+str)
		if err != nil {
			fmt.Println(err)
			return
		}
		SpidePage(img, movieNum)
	}
}

func SpidePage(img, movieNum string) {
	SaveJokeFile(img, movieNum)
}
