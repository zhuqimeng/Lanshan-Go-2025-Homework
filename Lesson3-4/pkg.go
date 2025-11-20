package main

import (
	"fmt"
	"strconv"
)

var (
	restaurant string
	money      int
	day        = 1
	mode       int32
	MyKitchen  kitchen
	MyStore    storehouse
	MyMarket   market
)

type Guest struct {
	name, hello, taste, sauce, dish string
	pay, sex                        int
}

func NewGuest() Guest {
	return guestList[GetRand(len(guestList))]
}

func (g Guest) ShowUp() {
	str := "【" + g.name + "】：" + fmt.Sprintf("%q", g.hello)
	fmt.Println("有客人来到了你的餐厅：\n", str)
	var ask string
	if g.sex == 1 {
		ask = "他"
	} else if g.sex == 0 {
		ask = "她"
	} else {
		ask = "它"
	}
	if g.dish != "无" {
		ask += "想要一份"
		if g.taste != "无" {
			ask += g.taste + "味的"
		}
		if g.sauce != "无" {
			ask += "加" + g.sauce + "的"
		}
		ask += g.dish + "，并将支付" + strconv.Itoa(g.pay) + "元。"
	} else {
		ask += "只是来" + restaurant + "看了看，并摇摇头走开了。或许每个人喜欢的餐厅并不一样……"
	}
	fmt.Println(ask)
}

type kitchen struct {
	level, tidy int
}

type Kitchen interface {
	Upgrade()
	Cook(Guest)
}

func (k kitchen) Upgrade() {
	fmt.Println("提升厨房等级可以让你拥有更高的工作效率，也就是说，你能够在中午接待更多客人！")
	fmt.Printf("你现在的厨房等级为 %d，提升到下一阶段需要 %d 元。是否提升？(y/n)\n", MyKitchen.level, MyKitchen.level*100)
	op := GetChoice()
	switch op {
	case 'y':
		if money < MyKitchen.level*100 {
			fmt.Println("你没有足够的钱。")
		} else {
			money -= MyKitchen.level * 100
			MyKitchen.level++
			fmt.Printf("升级成功！你现在中午可以接待 %d 名顾客了。\n", MyKitchen.level)
		}
	case 'n':
		fmt.Println("留得青山在，不怕没柴烧————稳健的选择。")
	}
}

func (k kitchen) Cook(g Guest) {
	if g.pay == -1 {
		return
	}
	var id1, id2, id3 = -1, -1, -1
	if g.taste != "无" {
		for i, v := range MyStore.taste {
			if v.name == g.taste && v.number > 0 {
				id1 = i
				break
			}
		}
		if id1 == -1 {
			fmt.Println("由于你的仓库里缺少", g.taste, "，无法做出这道菜。")
			return
		}
	}
	if g.sauce != "无" {
		for i, v := range MyStore.sauce {
			if v.name == g.sauce && v.number > 0 {
				id2 = i
				break
			}
		}
		if id2 == -1 {
			fmt.Println("由于你的仓库里缺少", g.sauce, "，无法做出这道菜。")
			return
		}
	}
	if g.dish != "无" {
		for i, v := range MyStore.dish {
			if v.name == g.dish && v.number > 0 {
				id3 = i
				break
			}
		}
		if id3 == -1 {
			fmt.Println("由于你的仓库里缺少", g.dish, "，无法做出这道菜。")
			return
		}
	}
	fmt.Println("你是否愿意售卖这道菜？(y/n)")
	op := GetChoice()
	switch op {
	case 'y':
		fmt.Println("【你】：用心提供每一道菜是本店的宗旨。")
		if id1 != -1 {
			MyStore.taste[id1].number--
		}
		if id2 != -1 {
			MyStore.sauce[id2].number--
		}
		MyStore.dish[id3].number--
		chk := GetRand(100)
		if chk > MyKitchen.tidy {
			fmt.Println("由于厨房糟糕的卫生，顾客拒绝为你的菜品付款。（是时候打扫一下厨房了！）")
		} else {
			money += g.pay
			MyKitchen.tidy -= 5
			fmt.Println("你获得了", g.pay, "元。")
		}
	case 'n':
		fmt.Println("【你】：商人不做亏本的买卖。")
		fmt.Println("顾客遗憾地离开了。")
	}
}

type stuff struct {
	name            string
	endtime, number int
}

type Stuff interface {
	Check()
}

func (s stuff) Check() {
	fmt.Printf("这是%d个%s，预计会在%d日晚上腐坏。（还剩%d天）\n", s.number, s.name, s.endtime, s.endtime-day)
}

type storehouse struct {
	taste, sauce, dish []stuff
}

type Storehouse interface {
	Add(byte, stuff)
	Clear()
	Check()
}

func (s *storehouse) Add(ch byte, one stuff) {
	if ch == 't' {
		s.taste = append(s.taste, one)
	} else if ch == 's' {
		s.sauce = append(s.sauce, one)
	} else if ch == 'd' {
		s.dish = append(s.dish, one)
	} else {
		panic("invalid ch")
	}
}

func (s *storehouse) Clear() {
	s.taste = DeleteSlice(s.taste)
	s.sauce = DeleteSlice(s.sauce)
	s.dish = DeleteSlice(s.dish)
}

func (s *storehouse) Check() {
	l1, l2, l3 := len(s.taste), len(s.sauce), len(s.dish)
	fmt.Println("------------------------")
	fmt.Printf("仓库里共有%d件存货。 \n", l1+l2+l3)
	fmt.Printf("配菜共有%d件：\n", l1)
	for _, t := range s.taste {
		t.Check()
	}
	fmt.Printf("酱料共有%d件：\n", l2)
	for _, t := range s.sauce {
		t.Check()
	}
	fmt.Printf("主食共有%d件：\n", l3)
	for _, d := range s.dish {
		d.Check()
	}
	fmt.Println("----")
}

type market struct {
	gala int
}

func (m market) Welcome() {
	fmt.Println("【前台】：欢迎光临食材商店，您有什么需要的吗？")
	fmt.Println("0.我就看看")
	for id := 0; id < 9; id++ {
		fmt.Printf("%d.新鲜的 %s，价格为 %d\n", id+1, foods[id], price[id])
	}
	var op, val, cnt int
	_, err := fmt.Scanf("%d\n", &op)
	if err != nil {
		panic(err)
	}
	if op < 1 || op > 9 {
		val = 0
		goto end
	}
	op--
	fmt.Printf("【前台】：好的，%s，您需要多少份？\n", foods[op])
	_, err = fmt.Scanf("%d\n", &cnt)
	if err != nil {
		panic(err)
	}
	if cnt < 0 {
		panic("expect a positive integer")
	}
	val = price[op] * cnt
	if money < val {
		goto end
	}
	switch op {
	case 0, 1, 2:
		MyStore.Add('t', stuff{
			name:    foods[op],
			endtime: day + 7,
			number:  cnt,
		})
	case 3, 4, 5:
		MyStore.Add('s', stuff{
			name:    foods[op],
			endtime: day + 7,
			number:  cnt,
		})
	case 6, 7, 8:
		MyStore.Add('d', stuff{
			name:    foods[op],
			endtime: day + 7,
			number:  cnt,
		})
	}
end:
	if m.gala < day {
		m.gala = GetRand(15) + day
	}
	if m.gala == day {
		fmt.Println("【前台】：恭喜您，本商店今日大促销，全场七折！")
		val = val * 10 / 7
		m.gala = GetRand(15) + day + 1
	} else {
		fmt.Printf("【前台】：对了，顺带一提，本商店将在第 %d 日进行大促销，千万不要错过哦！\n", m.gala)
	}
	if val == 0 || val > money {
		if val > money {
			fmt.Println("【你】（小声）：预算好像不够啊……")
		}
		fmt.Println("你什么也没买，两手空空地离开了商店。")
		return
	}
	money -= val
	fmt.Printf("你支付了 %d 元，并离开了商店。\n", val)
}

type Achievement struct {
	fg       bool
	check    func()
	describe string
}
