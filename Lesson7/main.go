package main

import (
	"Lesson07/api"
	"Lesson07/dao"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:your_sql_password@tcp(127.0.0.1:3306)/testbase?charset=utf8mb4&parseTime=True&loc=Local"
	myDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败: ", err)
	}
	dao.InitDB(myDb)
	api.InitRouterGin()
}
