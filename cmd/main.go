package main

import (
	"JD/api"
	"JD/dao"
)

func main() {
	dao.InitDB()
	api.InitEngine()
}
