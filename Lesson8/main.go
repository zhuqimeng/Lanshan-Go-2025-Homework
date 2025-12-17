package main

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/redis/go-redis/v9"
)

type Role struct {
	Role   string
	Gender string
	Age    string
}

func main() {
	cli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	cli.HSet(ctx, "魔裁", "樱羽艾玛", "粉色小狗")

	res := cli.HGet(ctx, "魔裁", "樱羽艾玛")

	fmt.Println(res.Val())

	cli.HDel(ctx, "魔裁", "樱羽艾玛")

	Role := Role{
		Role:   "樱羽艾玛",
		Gender: "粉色小狗",
		Age:    "18",
	}

	/*cli.HSet(ctx, "樱羽艾玛", map[string]interface{}{
		"Role":   Role.Role,
		"Gender": Role.Gender,
		"Age":    Role.Age,
	})*/

	cli.HSet(ctx, "data", map[string]string{
		"Role":   "np",
		"Gender": Role.Gender,
		"Age":    Role.Age,
	})

	re := cli.HGetAll(ctx, "data")

	mp, err := re.Result()
	if err != nil {
		fmt.Println(err)
	}
	// 一个将 map 对象反序列化为结构体的工具包
	err = mapstructure.Decode(mp, &Role)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", Role)
}
