package dao

import (
	"Lesson07/model"
	"Lesson07/utils"
	"Lesson07/utils/bf"
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	Name     string `gorm:"size:64;not null;comment:代办项目"`
	Priority string `gorm:"size:32;default:低;comment:优先级"`
	Status   string `gorm:"size:32;default:未开始;comment:任务状态"`
	Version  uint   `gorm:"default:0;comment:版本号机制"`
}

var (
	db     *gorm.DB
	cli    *redis.Client
	ctx    = context.Background()
	userBf *bf.BloomFilter
	dataBf *bf.BloomFilter
)

func GetToDos(req *model.QueryRequest) (*ToDo, error) {
	if req.Name == "" {
		return nil, errors.New("请输入要查询的项目名称")
	}
	if exists, err := dataBf.Exists(ctx, req.Name); err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("不存在的项目名称")
	}
	// 使用布隆过滤器防止缓存穿透
	isMember, err := cli.Exists(ctx, "Data:"+req.Name).Result()
	if err != nil {
		return nil, err
	}
	if isMember == 1 {
		res, err := cli.HGetAll(ctx, "Data:"+req.Name).Result()
		if err != nil {
			return nil, err
		}
		version, _ := strconv.Atoi(res["Version"])
		data := &ToDo{
			Name:     req.Name,
			Priority: res["Priority"],
			Status:   res["Status"],
			Version:  uint(version),
		}
		return data, nil
	}
	// 优先使用 redis - map 进行缓存查询
	var tmp ToDo
	res := db.Model(&ToDo{}).Where("name = ?", req.Name).First(&tmp)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	cli.HSet(ctx, "Data:"+req.Name, map[string]interface{}{
		"Name":     tmp.Name,
		"Priority": tmp.Priority,
		"Status":   tmp.Status,
		"Version":  tmp.Version,
	}, time.Duration(20+utils.GetRand(1, 5))*time.Minute)
	return &tmp, nil
}

func AddToDo(req *model.QueryRequest) error {
	if req.Name == "" {
		return errors.New("代办项目名称不能为空")
	}
	if exists, _ := cli.Exists(ctx, "Data:"+req.Name).Result(); exists == 1 {
		return errors.New("已存在的项目名称")
	}
	var count int64
	if err := db.Model(&ToDo{}).Where("name = ?", req.Name).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("已存在的项目名称")
	}
	// 保证项目名称不重复
	tmp := &ToDo{
		Name:     req.Name,
		Status:   req.Status,
		Priority: req.Priority,
	}
	result := db.Create(tmp)
	if result.Error != nil {
		return result.Error
	}
	cli.HSet(ctx, "Data:"+req.Name, map[string]interface{}{
		"Name":     tmp.Name,
		"Status":   tmp.Status,
		"Priority": tmp.Priority,
		"Version":  tmp.Version,
	}, time.Duration(20+utils.GetRand(1, 5))*time.Minute)
	// 建立新的缓存
	err := dataBf.Add(ctx, req.Name)
	return err
}

func UpdateToDo(req *model.QueryRequest) error {
	if exists, err := dataBf.Exists(ctx, req.Name); err != nil {
		return err
	} else if !exists {
		return errors.New("不存在的项目名称")
	}
	// 直接用 mySQL 进行写入操作，然后更新 redis 版本号
	var tmp ToDo
	res := db.Model(&ToDo{}).Where("name = ?", req.Name).First(&tmp)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	if req.Status != "" {
		tmp.Status = req.Status
	}
	if req.Priority != "" {
		tmp.Priority = req.Priority
	}
	tmp.Version++
	if err := db.Save(&tmp).Error; err != nil {
		return err
	}
	if ok, _ := cli.Exists(ctx, "Data:"+req.Name).Result(); ok == 1 {
		res, err := cli.HGetAll(ctx, "Data:"+req.Name).Result()
		if err != nil {
			return err
		}
		version, _ := strconv.Atoi(res["Version"])
		data := &ToDo{
			Name:     req.Name,
			Priority: res["Priority"],
			Status:   res["Status"],
			Version:  uint(version),
		}
		if data.Version >= tmp.Version {
			return nil
		} // 旧版本不进行更新
	}
	cli.HSet(ctx, "Data", req.Name, map[string]interface{}{
		"Name":     tmp.Name,
		"Status":   tmp.Status,
		"Priority": tmp.Priority,
		"Version":  tmp.Version,
	}, time.Duration(20+utils.GetRand(1, 5))*time.Minute)
	return nil
}

func DeleteToDo(req *model.QueryRequest) error {
	if exists, err := dataBf.Exists(ctx, req.Name); err != nil {
		return err
	} else if !exists {
		return errors.New("不存在的项目名称")
	}
	res := db.Model(&ToDo{}).Where("name = ?", req.Name).Delete(&ToDo{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("记录不存在，删除失败")
	}
	_, err := cli.Del(ctx, "Data:"+req.Name).Result()
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(req *model.CreateUserRequest) error {
	// 检查用户名是否已存在
	isMember, _ := cli.HExists(ctx, "User:", req.Username).Result()
	if isMember == true {
		return errors.New("已存在的用户名")
	}
	// 优先使用 redis - map 进行缓存查询
	var count int64
	if err := db.Model(&model.User{}).Where("username = ?", req.Username).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("已存在的用户名")
	}

	// 创建用户
	user := &model.User{
		Username: req.Username,
		Password: utils.HashPassword(req.Password), // 密码需要哈希
	}
	if err := db.Create(user).Error; err != nil {
		return err
	}
	cli.HSet(ctx, "User:", user.Username, user.Password, time.Duration(20+utils.GetRand(1, 5))*time.Minute)
	// 将用户信息写入缓存
	err := userBf.Add(ctx, req.Username)
	return err
}

func CheckUser(req *model.User) (string, error) {
	if exists, err := userBf.Exists(ctx, req.Username); err != nil {
		return "", err
	} else if !exists {
		return "", errors.New("不存在的用户名称")
	}
	// 使用布隆过滤器防止缓存穿透
	isMember, _ := cli.HExists(ctx, "User:", req.Username).Result()
	if isMember == true {
		password := cli.HGet(ctx, "User:", req.Username).Val()
		if utils.ComparePasswords(password, []byte(req.Password)) == false {
			return "", errors.New("错误的密码")
		}
		token, err := utils.MakeToken(req.Username, time.Now().Add(10*time.Minute))
		if err != nil {
			return "", err
		}
		return token, nil
	}
	// 优先使用 redis - map 进行缓存查询

	var user model.User
	res := db.Model(&model.User{}).Where("username = ?", req.Username).First(&user)
	if res.Error != nil {
		return "", res.Error
	}
	cli.HSet(ctx, "User:", user.Username, user.Password, time.Duration(20+utils.GetRand(1, 5))*time.Minute)
	// 不采用相同的过期时间来防止缓存雪崩。
	if utils.ComparePasswords(user.Password, []byte(req.Password)) == false {
		return "", errors.New("错误的密码")
	}
	token, err := utils.MakeToken(req.Username, time.Now().Add(10*time.Minute))
	if err != nil {
		return "", err
	}
	return token, nil
}

func InitDB(myDb *gorm.DB, myCli *redis.Client) {
	db = myDb
	cli = myCli
	err := db.AutoMigrate(&ToDo{}, &model.User{})
	if err != nil {
		log.Fatal("自动迁移失败: ", err)
	}
	log.Println("自动迁移成功")
	userBf = bf.NewBloomFilter(cli, "bloom:User", 10000, 0.001)
	dataBf = bf.NewBloomFilter(cli, "bloom:Data", 10000, 0.001)
}
