package main

import (
	"JD/api"
	"JD/dao"
)

func main() {
	api.InitEngine()
	dao.InitDB()
}
