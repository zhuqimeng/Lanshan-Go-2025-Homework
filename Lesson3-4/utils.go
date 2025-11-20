package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func DeleteSlice(slice []stuff) []stuff {
	var newSlice []stuff
	for _, s := range slice {
		if s.endtime == day || s.number == 0 {
			continue
		}
		newSlice = append(newSlice, s)
	}
	return newSlice
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRand(limit int) int {
	res := r.Intn(limit)
	return res
}

func GetChoice() byte {
	var res byte
	_, err := fmt.Scanf("%c\n", &res)
	if err != nil {
		panic(err)
	}
	if res != 'y' && res != 'n' {
		panic("expect 'y' or 'n'")
	}
	return res
}

func Morning() {
	fmt.Printf("%s经营的第 %d 天，%s\n", restaurant, day, hitokoto[GetRand(len(hitokoto))])
	fmt.Println(hintMorning)
	var op int
	_, err := fmt.Scanf("%d\n", &op)
	if err != nil {
		panic(err)
	}
	if op == 1 {
		MyMarket.Welcome()
	} else if op == 2 {
		MyKitchen.tidy = 100
		fmt.Println("你花了一上午进行大扫除，现在厨房的每个角落都一尘不染！")
	} else if op == 3 {
		MyKitchen.Upgrade()
	} else {
		zeroCnt++
	}
	goStep()
}

func Noon() {
	fmt.Println(hintNoon)
	for i := 0; i < MyKitchen.level; i++ {
		g := NewGuest()
		g.ShowUp()
		MyKitchen.Cook(g)
		goStep()
	}
	fmt.Println("中午结束了，睡个午觉吧……")
	goStep()
}

func Evening() {
	fmt.Println(hintEvening)
	MyStore.Clear()
	var op int
	_, err := fmt.Scanf("%d\n", &op)
	if err != nil {
		panic(err)
	}
	if op == 1 {
		MyStore.Check()
	} else if op == 2 {
		fmt.Printf("你的小金库里还剩下 %d 元。\n", money)
	} else if op == 3 {
		fmt.Println("您放弃了游戏。")
		os.Exit(0)
	} else {
		zeroCnt++
	}
	goStep()
}

func goStep() {
	fmt.Println("请按回车键继续...")
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return
	}
}

func Init() {
	fmt.Println("请输入你的餐厅名字：")
	_, err := fmt.Scan(&restaurant)
	if err != nil {
		panic(err)
	}
	fg := HasSuffix(restaurant, "餐厅")
	if !fg {
		restaurant = restaurant + "餐厅"
	}
	fmt.Printf("恭喜，%s今天就正式成立了，请妥善经营哦！\n", restaurant)
	fmt.Println(hintMode)
	_, err = fmt.Scanf("%d", &mode)
	if err != nil {
		panic(err)
	}
	switch mode {
	case 1:
		money = 100
		MyKitchen = kitchen{
			level: 1,
			tidy:  100,
		}
		for id, v := range foods {
			tmp := stuff{
				name:    v,
				endtime: 7,
				number:  2,
			}
			if id <= 2 {
				MyStore.Add('t', tmp)
			} else if id <= 5 {
				MyStore.Add('s', tmp)
			} else if id <= 8 {
				MyStore.Add('d', tmp)
			}
		}
		MyMarket.gala = GetRand(15) + day
	case 2:
		money = 2000
		fmt.Println("敬请期待……")
		os.Exit(0)
	case 3:
		fmt.Println("敬请期待……")
		os.Exit(0)
	default:
		panic("expected integer 1-3")
	}
	fmt.Println("那么，游戏开始……")
}

func BreakJudge() bool {
	if money < 0 {
		fmt.Println("很遗憾！", restaurant, "最终因资金链断裂破产了……游戏结束。")
		return false
	}
	if day > 999 {
		fmt.Println("三年之期已到，游戏自动结束，感谢支持！")
		return false
	}
	return true
}

func CheckAchieve() {
	for id, v := range achievementList {
		if v.fg == false && v.check() {
			fmt.Println("你解锁了成就，", v.describe)
			achievementList[id].fg = true
		}
	}
	var lose int
	lose = min(10+(day/10), 30)
	money -= lose
	fmt.Printf("凌晨 00：00，自动支付了餐厅门面租金 %d 元。\n", lose)
	goStep()
}

func Summary() {
	fmt.Printf("下面是游戏总结：\n%s最终赚了 %d 元。", restaurant, money)
	fmt.Println("这是你获得的成就：")
	for _, v := range achievementList {
		if v.fg == true {
			fmt.Println(v.describe)
		}
	}
}
