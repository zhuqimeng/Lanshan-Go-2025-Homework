package main

var guestList = []Guest{
	{
		name:  "老张",
		hello: "小伙子，做生意嘛，最讲究的就是诚信……不要让我发现你偷工减料啊！",
		taste: "蔬菜",
		sauce: "番茄酱",
		dish:  "面条",
		pay:   15,
		sex:   1,
	},
	{
		name:  "老张",
		hello: "诚信是做人之本，经商也是一样。",
		taste: "水果",
		sauce: "沙拉酱",
		dish:  "米饭",
		pay:   18,
		sex:   1,
	},
	{
		name:  "多啦A梦",
		hello: "心情好的话什么都能顺利，垂头丧气才会招来不幸。",
		taste: "牛肉",
		sauce: "无",
		dish:  "烤鱼",
		pay:   40,
		sex:   -1,
	},
	{
		name:  "多啦A梦",
		hello: "之前一直觉得很平常的事，到了现在才觉得如此重要……",
		taste: "蔬菜",
		sauce: "辣椒酱",
		dish:  "米饭",
		pay:   32,
		sex:   -1,
	},
	{
		name:  "朵拉",
		hello: "你看见我的帽子了吗？",
		taste: "无",
		sauce: "无",
		dish:  "米饭",
		pay:   10,
		sex:   0,
	},
	{
		name:  "朵拉",
		hello: "这是什么地方？",
		taste: "无",
		sauce: "无",
		dish:  "无",
		pay:   -1,
		sex:   0,
	},
	{
		name:  "初音未来",
		hello: "日暮江水远，入夜随风迁。夜未香未眠，寻花情已倦。",
		taste: "无",
		sauce: "沙拉酱",
		dish:  "烤鱼",
		pay:   25,
		sex:   0,
	},
	{
		name:  "初音未来",
		hello: "我的故乡没有蔷薇海，却有和你的双眸一样蔚蓝的天空。",
		taste: "牛肉",
		sauce: "辣椒酱",
		dish:  "面条",
		pay:   17,
		sex:   0,
	},
}

var achievementList = []Achievement{
	{
		fg: false,
		check: func() bool {
			if money < 0 {
				return true
			}
			return false
		},
		describe: "【倾家荡产】：你花光了所有积蓄。",
	},
	{
		fg: false,
		check: func() bool {
			if money > 300 {
				return true
			}
			return false
		},
		describe: "【小有所获】：第一次存款达到 300 元。",
	},
	{
		fg: false,
		check: func() bool {
			if serveCnt > 0 {
				return true
			}
			return false
		},
		describe: "【第一桶金】：第一次为顾客提供菜品。",
	},
	{
		fg: false,
		check: func() bool {
			if refuseCnt > 0 {
				return true
			}
			return false
		},
		describe: "【学会拒绝】：第一次拒绝为顾客上菜。",
	},
	{
		fg: false,
		check: func() bool {
			if badCnt > 0 {
				return true
			}
			return false
		},
		describe: "【邋里邋遢】：第一次因为厨房太脏而被吃霸王餐。",
	},
	{
		fg: false,
		check: func() bool {
			if payNumber > 0 {
				return true
			}
			return false
		},
		describe: "【及时补充】：第一次去商店补充食材。",
	},
}
