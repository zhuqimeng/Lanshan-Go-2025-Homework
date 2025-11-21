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

func GetChoice() string {
	var res string
	for {
		_, err := fmt.Scan(&res)
		if res != "y" && res != "n" {
			continue
		}
		if err == nil {
			break
		}
	}
	return res
}

func smallGame() {
	fmt.Println("【庄家】：输入你要押注的钱，然后投掷骰子……有一半的概率获得双倍，一半的概率失去所有。")
	var x int
	for {
		_, err := fmt.Scanln(&x)
		if x <= 0 {
			fmt.Println("你需要至少押注一元。")
			continue
		}
		if x > money {
			fmt.Printf("你只有 %d 元，请重新输入。\n", money)
			continue
		}
		if err == nil {
			break
		}
	}
	chk := GetRand(100) + 1
	fmt.Println("你骰出了点数：", chk)
	if chk > 50 {
		fmt.Println("【庄家】：小子，运气不错啊，拿好你的钱。")
		money += x
		winMoney += x
		fmt.Println("你获得了", x, "元。")
	} else {
		fmt.Println("【庄家】：我这里可没有后悔药卖，好好开你的餐厅吧……")
		money -= x
		lostMoney += x
		fmt.Println("你失去了", x, "元。")
	}
}

func Morning() {
	fmt.Printf("%s经营的第 %d 天，%s\n", restaurant, day, hitokoto[GetRand(len(hitokoto))])
	fmt.Println(hintMorning)
	var op int
	for {
		_, err := fmt.Scanln(&op)
		if err == nil {
			break
		}
	}
	if op == 1 {
		MyMarket.Welcome()
	} else if op == 2 {
		MyKitchen.tidy = 100
		fmt.Println("你花了一上午进行大扫除，现在厨房的每个角落都一尘不染！")
	} else if op == 3 {
		MyKitchen.Upgrade()
	} else if op == 4 {
		smallGame()
	} else {
		zeroCnt++
	}
	goStep()
}

func Noon() {
	fmt.Println(hintNoon)
	dayIncome = 0
	if zeroCnt > 0 && zeroCnt%5 == 0 {
		fmt.Println("不知为何，你觉得今天头脑清醒，手脚灵活，干起活来更有效率了……")
		g := NewGuest()
		g.ShowUp()
		MyKitchen.Cook(g)
		goStep()
	}
	for i := 0; i < MyKitchen.level; i++ {
		g := NewGuest()
		g.ShowUp()
		MyKitchen.Cook(g)
		goStep()
	}
	fmt.Println("中午结束了，睡个午觉吧……")
	goStep()
}

func Evening() bool {
	fmt.Println(hintEvening)
	MyStore.Clear()
	var op int
	for {
		_, err := fmt.Scanln(&op)
		if err == nil {
			break
		}
	}
	if op == 1 {
		MyMarket.Welcome()
	} else if op == 2 {
		MyStore.Check()
	} else if op == 3 {
		fmt.Printf("你的小金库里还剩下 %d 元。\n", money)
	} else if op == 4 {
		fmt.Println("你想了想，觉得开餐厅这件事还是不太适合自己……")
		fmt.Println("于是你决定在今天结束时将", restaurant, "转让给其他人。")
		goStep()
		return true
	} else {
		zeroCnt++
	}
	goStep()
	return false
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
	for {
		_, err := fmt.Scanln(&restaurant)
		if err == nil {
			break
		}
	}

	fg := HasSuffix(restaurant, "餐厅")
	if !fg {
		restaurant = restaurant + "餐厅"
	}
	fmt.Printf("恭喜，%s今天就正式成立了，请妥善经营哦！\n", restaurant)
	fmt.Println(hintMode)
	for {
		_, err := fmt.Scanln(&mode)
		if err == nil {
			break
		}
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
				endtime: 10,
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
		return true
	}
	if day > 999 {
		fmt.Println("三年之期已到，游戏自动结束，感谢支持！")
		return true
	}
	return false
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
	fmt.Println(gap1)
	fmt.Printf("下面是游戏总结：\n%s经营了 %d 天，最终赚了 %d 元。\n", restaurant, day, money)
	fmt.Println("这是你获得的成就：")
	for _, v := range achievementList {
		if v.fg == true {
			fmt.Println(v.describe)
		}
	}
	fmt.Println(gap2)
	goStep()
}
