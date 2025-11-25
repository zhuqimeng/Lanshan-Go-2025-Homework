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
		if res != "y" && res != "n" && res != "yes" && res != "no" && res != "YES" && res != "NO" {
			continue
		}
		if err == nil {
			break
		}
		fmt.Println("请再次输入。")
	}
	if res == "YES" || res == "yes" {
		res = "y"
	}
	if res == "NO" || res == "no" {
		res = "n"
	}
	return res
}

func smallGame() {
	if money == 0 {
		fmt.Println("你现在没有钱进行押注。")
		return
	}
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
		fmt.Println("请再次输入。")
	}
	maxStake = max(maxStake, x)
	chk := GetRand(100) + 1
	fmt.Println("你骰出了点数：", chk)
	if chk > 50 {
		fmt.Println("【庄家】：小子，运气不错啊，拿好你的钱。")
		money += x
		winMoney += x
		fmt.Println("你获得了", x, "元。")
	} else {
		fmt.Println("【庄家】：我这里可没有后悔药卖，好好开你的餐厅吧……")
		if money == x {
			allPut = true
		}
		money -= x
		lostMoney += x
		fmt.Println("你失去了", x, "元。")
	}
}

func Morning() {
	fmt.Printf("%s经营的第 %d 天，%s\n", restaurant, day, hitokoto[GetRand(len(hitokoto))])
	fmt.Println(hintMorning)
Loop:
	var op int
	for {
		_, err := fmt.Scanln(&op)
		if err == nil {
			break
		}
		fmt.Println("请再次输入。")
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
	} else if op == 0 {
		zeroCnt++
	} else {
		fmt.Println("请再次输入。")
		goto Loop
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

func Afternoon() {
	fmt.Println(hintAfternoon)
	var op int
Loop:
	for {
		_, err := fmt.Scanln(&op)
		if err == nil {
			break
		}
		fmt.Println("请再次输入。")
	}
	if op == 1 {
		MyStore.Freezing()
	} else if op == 2 {
		earn := zeroCnt*7 + GetRand(20) + 1
		money += earn
		maxEarn = max(earn, maxEarn)
		zeroCnt = 0
		fmt.Println("恭喜！你通过兼职一下午换取了", earn, "元。")
	} else if op == 0 {
		zeroCnt++
	} else {
		fmt.Println("请再次输入。")
		goto Loop
	}
	goStep()
}

func Evening() bool {
	fmt.Println(hintEvening)
	MyStore.Clear()
Loop:
	var op int
	for {
		_, err := fmt.Scanln(&op)
		if err == nil {
			break
		}
		fmt.Println("请再次输入。")
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
	} else if op == 0 {
		zeroCnt++
	} else {
		fmt.Println("请再次输入。")
		goto Loop
	}
	goStep()
	return false
}

func goStep() {
	fmt.Println("请按回车键继续...")
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return
	}
}

func Init() {
	fmt.Println("请输入你的餐厅名字：（请不要输入带空格的名字）")
	for {
		_, err := fmt.Scanln(&restaurant)
		if err == nil {
			break
		}
		fmt.Println("请再次输入。")
	}
	fg := HasSuffix(restaurant, "餐厅")
	if !fg {
		restaurant = restaurant + "餐厅"
	}
	fmt.Printf("恭喜，%s今天就正式成立了，请妥善经营哦！\n", restaurant)
	fmt.Println(hintMode)
Loop:
	for {
		_, err := fmt.Scanln(&mode)
		if err == nil {
			break
		}
		fmt.Println("请再次输入。")
	}
	switch mode {
	case 1:
		money = 150
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
		goStep()
		os.Exit(0)
	case 3:
		fmt.Println("敬请期待……")
		goStep()
		os.Exit(0)
	default:
		fmt.Println("请再次输入。")
		goto Loop
	}
	fmt.Println("你获得了启动资金", money, "元。")
	fmt.Println("那么，游戏开始……")
	fmt.Println("通关条件：经营餐馆一百天或者存款达到一万元。")
	goStep()
}

func BreakJudge() bool {
	if money < 0 {
		fmt.Printf("\033[1;31;40m%s %s %s\033[0m\n", "很遗憾！", restaurant, "最终因资金链断裂破产了……游戏结束。")
		return true
	}
	if okFlag == false && (day >= 100 || money >= 10000) {
		fmt.Println("恭喜！你通关了游戏。感谢支持。")
		fmt.Println("是否继续游戏？(y/n)")
		op := GetChoice()
		fmt.Println()
		if op == "y" {
			okFlag = true
			return false
		}
		return true
	}
	return false
}

func CheckAchieve() {
	if MyMarket.gala == day {
		MyMarket.gala = GetRand(15) + day + 1
	}
	for id, v := range achievementList {
		if v.fg == false && v.check() {
			fmt.Println("你解锁了成就，", v.describe)
			achievementList[id].fg = true
			money += v.bonus
			fmt.Println("(奖励", v.bonus, "元。)")
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
