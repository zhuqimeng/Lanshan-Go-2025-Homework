实际上，我在实现版本号更新时使用了分布式锁，那么这其实时悲观锁策略。

不过我的初衷是使用乐观锁，因此这里贴一段使用乐观锁实现的代码：

```go
func UpdateToDoOptimistic(req *model.QueryRequest) error {
    // 使用布隆过滤器检查项目是否存在
    if exists, err := dataBf.Exists(ctx, req.Name); err != nil {
        return err
    } else if !exists {
        return errors.New("不存在的项目名称")
    }

    maxRetries := 3
    for retry := 0; retry < maxRetries; retry++ {
        // 1. 获取当前版本
        var tmp ToDo
        res := db.Model(&ToDo{}).Where("name = ?", req.Name).First(&tmp)
        if errors.Is(res.Error, gorm.ErrRecordNotFound) {
            return res.Error
        }
        
        currentVersion := tmp.Version
        
        // 2. 准备更新
        if req.Status != "" {
            tmp.Status = req.Status
        }
        if req.Priority != "" {
            tmp.Priority = req.Priority
        }
        tmp.Version++  // 版本号增加
        
        // 3. 乐观锁更新：使用CAS（Compare And Swap）
        // 只在版本号匹配时更新
        result := db.Model(&ToDo{}).
            Where("name = ? AND version = ?", req.Name, currentVersion).
            Updates(map[string]interface{}{
                "status":   tmp.Status,
                "priority": tmp.Priority,
                "version":  tmp.Version,
            })
        
        if result.Error != nil {
            return result.Error
        }
        
        // 4. 检查是否更新成功（通过受影响的行数）
        if result.RowsAffected == 0 {
            // 版本冲突，重试
            if retry < maxRetries-1 {
                time.Sleep(time.Millisecond * time.Duration(10*(retry+1)))
                continue
            }
            return errors.New("更新失败，版本冲突")
        }
        
        // 5. 更新Redis（也需要版本检查）
        redisKey := "Data:" + req.Name
        // 使用Redis事务/WATCH确保原子性
        err := cli.Watch(ctx, func(tx *redis.Tx) error {
            // 检查Redis版本
            redisVersionStr, err := tx.HGet(ctx, redisKey, "Version").Result()
            if err != nil && err != redis.Nil {
                return err
            }
            
            // 如果Redis版本 >= 当前版本，不更新
            if redisVersionStr != "" {
                redisVersion, _ := strconv.ParseUint(redisVersionStr, 10, 32)
                if uint(redisVersion) >= tmp.Version {
                    return nil // 跳过更新
                }
            }
            
            // 更新Redis
            _, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
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
        }, redisKey)
        
        if err != nil {
            return err
        }
        
        return nil
    }
    
    return errors.New("更新失败，重试次数超限")
}
```

