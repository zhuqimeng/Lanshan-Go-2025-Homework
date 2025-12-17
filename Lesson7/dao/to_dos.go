package dao

import (
	"Lesson07/model"
	"Lesson07/utils"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func GetToDos(req *model.QueryRequest) (*model.ToDo, error) {
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
		data := &model.ToDo{
			Name:     req.Name,
			Priority: res["Priority"],
			Status:   res["Status"],
			Version:  uint(version),
		}
		return data, nil
	}
	// 优先使用 redis - map 进行缓存查询
	var tmp model.ToDo
	res := db.Model(&model.ToDo{}).Where("name = ?", req.Name).First(&tmp)
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
	// 创建分布式锁键
	lockKey := "lock:todo:" + req.Name
	lockValue := fmt.Sprintf("%d", time.Now().UnixNano())

	// 尝试获取锁，设置5秒过期时间
	locked, err := cli.SetNX(ctx, lockKey, lockValue, 5*time.Second).Result()
	if err != nil {
		return err
	}
	if !locked {
		return errors.New("系统繁忙，请稍后重试")
	}

	// 确保释放锁
	defer func() {
		// 使用Lua脚本确保只删除自己的锁
		luaScript := `
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("del", KEYS[1])
        else
            return 0
        end
        `
		cli.Eval(ctx, luaScript, []string{lockKey}, lockValue)
	}()
	var count int64
	if err := db.Model(&model.ToDo{}).Where("name = ?", req.Name).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("已存在的项目名称")
	}
	// 保证项目名称不重复
	tmp := &model.ToDo{
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
	err = dataBf.Add(ctx, req.Name)
	return err
}

func UpdateToDo(req *model.QueryRequest) error {
	// 使用布隆过滤器检查项目是否存在
	if exists, err := dataBf.Exists(ctx, req.Name); err != nil {
		return err
	} else if !exists {
		return errors.New("不存在的项目名称")
	}

	// 创建分布式锁键
	lockKey := "lock:todo:" + req.Name
	lockValue := fmt.Sprintf("%d", time.Now().UnixNano())

	// 尝试获取锁，设置5秒过期时间
	locked, err := cli.SetNX(ctx, lockKey, lockValue, 5*time.Second).Result()
	if err != nil {
		return err
	}
	if !locked {
		return errors.New("系统繁忙，请稍后重试")
	}

	// 确保释放锁
	defer func() {
		// 使用Lua脚本确保只删除自己的锁
		luaScript := `
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("del", KEYS[1])
        else
            return 0
        end
        `
		cli.Eval(ctx, luaScript, []string{lockKey}, lockValue)
	}()

	// 直接从MySQL获取数据
	var tmp model.ToDo
	res := db.Model(&model.ToDo{}).Where("name = ?", req.Name).First(&tmp)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}

	// 应用更新
	if req.Status != "" {
		tmp.Status = req.Status
	}
	if req.Priority != "" {
		tmp.Priority = req.Priority
	}
	tmp.Version++

	// 先更新MySQL
	if err := db.Save(&tmp).Error; err != nil {
		return err
	}

	// 再更新Redis（现在可以安全更新，因为有锁保护）
	redisKey := "Data:" + req.Name
	exists, err := cli.Exists(ctx, redisKey).Result()
	if err != nil {
		return err
	}

	if exists == 1 {
		// 获取Redis当前版本
		redisVersionStr, err := cli.HGet(ctx, redisKey, "Version").Result()
		if err != nil {
			return err
		}

		if redisVersionStr != "" {
			redisVersion, _ := strconv.ParseUint(redisVersionStr, 10, 32)
			// 如果Redis版本已经等于或高于MySQL版本，跳过更新
			if uint(redisVersion) >= tmp.Version {
				return nil
			}
		}
	}

	// 更新Redis
	_, err = cli.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, redisKey, map[string]interface{}{
			"Name":     tmp.Name,
			"Status":   tmp.Status,
			"Priority": tmp.Priority,
			"Version":  tmp.Version,
		})
		pipe.Expire(ctx, redisKey, time.Duration(20+utils.GetRand(1, 5))*time.Minute)
		return nil
	})

	return err
}

func DeleteToDo(req *model.QueryRequest) error {
	if exists, err := dataBf.Exists(ctx, req.Name); err != nil {
		return err
	} else if !exists {
		return errors.New("不存在的项目名称")
	}
	// 创建分布式锁键
	lockKey := "lock:todo:" + req.Name
	lockValue := fmt.Sprintf("%d", time.Now().UnixNano())

	// 尝试获取锁，设置5秒过期时间
	locked, err := cli.SetNX(ctx, lockKey, lockValue, 5*time.Second).Result()
	if err != nil {
		return err
	}
	if !locked {
		return errors.New("系统繁忙，请稍后重试")
	}

	// 确保释放锁
	defer func() {
		// 使用Lua脚本确保只删除自己的锁
		luaScript := `
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("del", KEYS[1])
        else
            return 0
        end
        `
		cli.Eval(ctx, luaScript, []string{lockKey}, lockValue)
	}()
	res := db.Model(&model.ToDo{}).Where("name = ?", req.Name).Delete(&model.ToDo{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("记录不存在，删除失败")
	}
	_, err = cli.Del(ctx, "Data:"+req.Name).Result()
	if err != nil {
		return err
	}
	return nil
}
