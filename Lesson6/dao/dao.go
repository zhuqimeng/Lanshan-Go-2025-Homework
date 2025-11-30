package dao

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// 模拟数据库
var database = map[string]string{}

func AddUser(username string, password string) {
	database[username] = password
}

func FindUser(username string, password string) bool {
	if pwd, ok := database[username]; ok {
		if pwd == password {
			return true
		}
	}
	return false
}

func UpdateUser(username string, password string) {
	database[username] = password
}

func Save() {
	file, err := os.OpenFile("dao/data.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(database)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Init() {
	file, err := os.Open("dao/data.json")
	if err != nil {
		if err == io.EOF {
			return
		}
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&database)
	if err != nil {
		if err == io.EOF {
			return
		}
		panic(err)
	}
}
