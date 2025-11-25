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
	{
		name:  "穿山甲",
		hello: "不好意思哈，走错片场了。",
		taste: "蔬菜",
		sauce: "无",
		dish:  "烤鱼",
		pay:   12,
		sex:   1,
	},
	{
		name:  "穿山甲",
		hello: "你这里有鸡汤卖吗？",
		taste: "无",
		sauce: "无",
		dish:  "无",
		pay:   -1,
		sex:   1,
	},
	{
		name:  "章鱼哥",
		hello: "你知道吗，竖笛是世界上最美妙的乐器~~~",
		taste: "水果",
		sauce: "无",
		dish:  "面条",
		pay:   14,
		sex:   -1,
	},
	{
		name:  "章鱼哥",
		hello: "艺术是要吃苦的。",
		taste: "蔬菜",
		sauce: "番茄酱",
		dish:  "面条",
		pay:   11,
		sex:   -1,
	},
	{
		name:  "爱莉希雅",
		hello: "亲爱的山雀，请将我的箭，我的花，与我的爱，带给那子然独行的旅人。",
		taste: "牛肉",
		sauce: "沙拉酱",
		dish:  "米饭",
		pay:   34,
		sex:   0,
	},
	{
		name:  "爱莉希雅",
		hello: "你好！新的一天，从一场美妙的邂逅开始。",
		taste: "无",
		sauce: "沙拉酱",
		dish:  "烤鱼",
		pay:   29,
		sex:   0,
	},
	{
		name:  "亚里士多德",
		hello: "我们日复一日做的事情，决定了我们是怎样的人。",
		taste: "牛肉",
		sauce: "无",
		dish:  "米饭",
		pay:   23,
		sex:   1,
	},
	{
		name:  "亚里士多德",
		hello: "幸福是把灵魂安放在最适当的位置。",
		taste: "无",
		sauce: "辣椒酱",
		dish:  "米饭",
		pay:   23,
		sex:   1,
	},
	{
		name:  "御坂美琴",
		hello: "别对自己说谎。",
		taste: "蔬菜",
		sauce: "辣椒酱",
		dish:  "烤鱼",
		pay:   45,
		sex:   0,
	},
	{
		name:  "神龙",
		hello: "遇到我是一种幸运。",
		taste: "牛肉",
		sauce: "沙拉酱",
		dish:  "烤鱼",
		pay:   77,
		sex:   -1,
	},
	{
		name:  "薇尔莉特",
		hello: "花无凋零之时，爱无传达之期，爱亘古不变，紫罗兰永世长存。",
		taste: "无",
		sauce: "沙拉酱",
		dish:  "米饭",
		pay:   19,
		sex:   0,
	},
	{
		name:  "薇尔莉特",
		hello: "我想，我们正是在伤害与被伤害的循环往复中，渐渐变得温柔起来的。",
		taste: "无",
		sauce: "沙拉酱",
		dish:  "面条",
		pay:   18,
		sex:   0,
	},
	{
		name:  "薇尔莉特",
		hello: "语言的含义有两层，说出口的并不是全部。",
		taste: "水果",
		sauce: "无",
		dish:  "烤鱼",
		pay:   21,
		sex:   0,
	},
}

var achievementList = []Achievement{
	{
		fg: false,
		check: func() bool {
			if money <= 0 {
				return true
			}
			return false
		},
		bonus:    6,
		describe: "【中道崩殂】：你花光了所有积蓄。",
	},
	{
		fg: false,
		check: func() bool {
			if money >= 500 {
				return true
			}
			return false
		},
		bonus:    50,
		describe: "【小有所获】：第一次存款达到 500 元。",
	},
	{
		fg: false,
		check: func() bool {
			if money >= 1000 {
				return true
			}
			return false
		},
		bonus:    100,
		describe: "【经济头脑】：第一次存款达到 1000 元。",
	},
	{
		fg: false,
		check: func() bool {
			if money >= 10000 {
				return true
			}
			return false
		},
		bonus:    666,
		describe: "【财务自由】：最终存款达到 10000 元。",
	},
	{
		fg: false,
		check: func() bool {
			if dayIncome >= 100 {
				return true
			}
			return false
		},
		bonus:    15,
		describe: "【日进斗金】：餐厅单日收入达到 100 元。",
	},
	{
		fg: false,
		check: func() bool {
			if serveCnt > 0 {
				return true
			}
			return false
		},
		bonus:    10,
		describe: "【第一桶金】：第一次为顾客提供菜品。",
	},
	{
		fg: false,
		check: func() bool {
			if serveCnt >= 10 {
				return true
			}
			return false
		},
		bonus:    30,
		describe: "【优质服务】：累计为顾客提供十次菜品。",
	},
	{
		fg: false,
		check: func() bool {
			if refuseCnt > 0 {
				return true
			}
			return false
		},
		bonus:    10,
		describe: "【学会拒绝】：第一次拒绝为顾客上菜。",
	},
	{
		fg: false,
		check: func() bool {
			if refuseCnt >= 10 {
				return true
			}
			return false
		},
		bonus:    30,
		describe: "【霸道老板】：累计十次拒绝为顾客提供菜品。",
	},
	{
		fg: false,
		check: func() bool {
			if badCnt > 0 {
				return true
			}
			return false
		},
		bonus:    15,
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
		bonus:    10,
		describe: "【及时补充】：第一次去商店补充食材。",
	},
	{
		fg: false,
		check: func() bool {
			if payNumber >= 100 {
				return true
			}
			return false
		},
		bonus:    30,
		describe: "【资金循环】：累计在商店消费 100 元。",
	},
	{
		fg: false,
		check: func() bool {
			if discount > 0 {
				return true
			}
			return false
		},
		bonus:    10,
		describe: "【捡个便宜】：在商店促销时购买一次食材。",
	},
	{
		fg: false,
		check: func() bool {
			if zeroCnt > 0 {
				return true
			}
			return false
		},
		bonus:    10,
		describe: "【小憩一下】：进行一次休息。",
	},
	{
		fg: false,
		check: func() bool {
			if zeroCnt >= 5 {
				return true
			}
			return false
		},
		bonus:    20,
		describe: "【精力充沛】：累计进行五次休息。",
	},
	{
		fg: false,
		check: func() bool {
			if zeroCnt >= 10 {
				return true
			}
			return false
		},
		bonus:    30,
		describe: "【养生流派】：累计进行十次休息。",
	},
	{
		fg: false,
		check: func() bool {
			if MyKitchen.level > 1 {
				return true
			}
			return false
		},
		bonus:    20,
		describe: "【休整一番】：对厨房进行一次升级。",
	},
	{
		fg: false,
		check: func() bool {
			if MyKitchen.level > 3 {
				return true
			}
			return false
		},
		bonus:    50,
		describe: "【节节高升】：对厨房进行三次升级。",
	},
	{
		fg: false,
		check: func() bool {
			if winMoney >= 300 {
				return true
			}
			return false
		},
		bonus:    20,
		describe: "【幸运骰子】：通过扔骰子累计赚取 300 元。",
	},
	{
		fg: false,
		check: func() bool {
			if lostMoney >= 300 {
				return true
			}
			return false
		},
		bonus:    20,
		describe: "【血本无归】：因为扔骰子累计失去 300 元。",
	},
	{
		fg: false,
		check: func() bool {
			if maxStake >= 1000 {
				return true
			}
			return false
		},
		bonus:    40,
		describe: "【一掷千金】：扔骰子单次押注金额超过 1000 元。",
	},
	{
		fg: false,
		check: func() bool {
			return allPut
		},
		bonus:    15,
		describe: "【倾家荡产】：因为扔骰子而输掉了所有存款。",
	},
	{
		fg: false,
		check: func() bool {
			if maxEarn >= 50 {
				return true
			}
			return false
		},
		bonus:    30,
		describe: "【兼职高手】：单日兼职赚取超过 50 元。",
	},
	{
		fg: false,
		check: func() bool {
			if day >= 10 {
				return true
			}
			return false
		},
		bonus:    10,
		describe: "【初出茅庐】：经营餐厅超过 10 天。",
	},
	{
		fg: false,
		check: func() bool {
			if day >= 30 {
				return true
			}
			return false
		},
		bonus:    30,
		describe: "【渐入佳境】：经营餐厅超过 30 天。",
	},
	{
		fg: false,
		check: func() bool {
			if day >= 70 {
				return true
			}
			return false
		},
		bonus:    70,
		describe: "【持之以恒】：经营餐厅超过 70 天。",
	},
	{
		fg: false,
		check: func() bool {
			if day >= 100 {
				return true
			}
			return false
		},
		bonus:    666,
		describe: "【商业巨亨】：经营餐厅超过 100 天。",
	},
}
