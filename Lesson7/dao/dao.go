package dao

import (
	"Lesson07/model"
	"Lesson07/utils"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	Name     string `gorm:"size:64;not null;comment:代办项目"`
	Priority string `gorm:"size:32;default:低;comment:优先级"`
	Status   string `gorm:"size:32;default:未开始;comment:任务状态"`
}

var db *gorm.DB

func GetToDos(req *model.QueryRequest) *gorm.DB {
	res := db.Model(&ToDo{})
	if req.Status != "" {
		res = res.Where("status = ?", req.Status)
	}
	if req.ID != 0 {
		res = res.Where("id = ?", req.ID)
	}
	if req.Name != "" {
		res = res.Where("name = ?", req.Name)
	}
	if req.Priority != "" {
		res = res.Where("priority = ?", req.Priority)
	}
	return res
}

func AddToDo(req *model.QueryRequest) error {
	if req.Name == "" {
		return errors.New("代办项目名称不能为空")
	}
	tmp := &ToDo{
		Name:     req.Name,
		Status:   req.Status,
		Priority: req.Priority,
	}
	result := db.Create(tmp)
	return result.Error
}

func UpdateToDo(req *model.QueryRequest) error {
	res := db.Model(&ToDo{}).Where("id = ?", req.ID)
	if res.Error != nil {
		return res.Error
	}
	var tmp ToDo
	if err := res.Find(&tmp).Error; err != nil {
		return err
	}
	if req.Name != "" {
		tmp.Name = req.Name
	}
	if req.Status != "" {
		tmp.Status = req.Status
	}
	if tmp.Priority != "" {
		tmp.Priority = req.Priority
	}
	if err := db.Save(&tmp).Error; err != nil {
		return err
	}
	return nil
}

func DeleteToDo(req *model.QueryRequest) error {
	res := db.Model(&ToDo{}).Where("id = ?", req.ID).Delete(&ToDo{})
	return res.Error
}

func CreateUser(req *model.CreateUserRequest) error {
	// 检查用户名是否已存在
	var count int64
	if err := db.Model(&model.User{}).Where("username = ?", req.Username).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户名已存在")
	}

	// 创建用户
	user := &model.User{
		Username: req.Username,
		Password: utils.HashPassword(req.Password), // 密码需要哈希
	}
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func CheckUser(req *model.User) (string, error) {
	var user model.User
	res := db.Model(&model.User{}).Where("username = ?", req.Username).First(&user)
	if res.Error != nil {
		return "", res.Error
	}
	if utils.ComparePasswords(user.Password, []byte(req.Password)) == false {
		return "", errors.New("错误的密码")
	}
	token, err := utils.MakeToken(req.Username, time.Now().Add(10*time.Minute))
	if err != nil {
		return "", err
	}
	user.Token = token
	if err = db.Save(&user).Error; err != nil {
		return "", err
	}
	return token, nil
}

func InitDB(myDb *gorm.DB) {
	db = myDb
	err := db.AutoMigrate(&ToDo{}, &model.User{})
	if err != nil {
		log.Fatal("自动迁移失败: ", err)
	}
	log.Println("自动迁移成功")
}
