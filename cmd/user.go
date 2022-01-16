package main

import (
	_ "github.com/go-sql-driver/mysql"
	"movie/api"
	"movie/dao"
)

func main() {
	dao.InitDB()
	api.InitEngine()
}
