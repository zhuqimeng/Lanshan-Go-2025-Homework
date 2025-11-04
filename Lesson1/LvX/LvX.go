package main

import "fmt"

func Average(sum, num int) float64 {
	res := float64(sum) / float64(num)
	return res
}

func main() {
	var x, sum, num int
	for {
		fmt.Println("请输入一个整数(输入0结束)：")
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		sum += x
		num++
	}
	var ans float64 = Average(sum, num)
	if ans >= 60 {
		fmt.Println("平均成绩为", ans, "，成绩合格")
	} else {
		fmt.Println("平均成绩为", ans, "，成绩不合格")
	}
}
