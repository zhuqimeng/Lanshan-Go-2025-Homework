package main

import (
	"Lesson07/api"
	"Lesson07/dao"
	"log"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Cyzhu8899312_@tcp(127.0.0.1:3306)/testbase?charset=utf8mb4&parseTime=True&loc=Local"
	myDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败: ", err)
	}
	cli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	dao.InitDB(myDb, cli)
	api.InitRouterGin()
}
