package main

func main() {
	Init()
	for ; BreakJudge(); day++ {
		Morning()
		Noon()
		Evening()
		CheckAchieve()
	}
	Summary()
}
