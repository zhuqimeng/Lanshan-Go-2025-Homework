package dao

import (
	"Lesson07/model"
	"Lesson07/utils"
	"context"
	"errors"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	cli    *redis.Client
	ctx    = context.Background()
	userBf *utils.BloomFilter
	dataBf *utils.BloomFilter
)

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
	err := db.AutoMigrate(&model.ToDo{}, &model.User{})
	if err != nil {
		log.Fatal("自动迁移失败: ", err)
	}
	log.Println("自动迁移成功")
	userBf = utils.NewBloomFilter(cli, "bloom:User", 10000, 0.001)
	dataBf = utils.NewBloomFilter(cli, "bloom:Data", 10000, 0.001)
}
