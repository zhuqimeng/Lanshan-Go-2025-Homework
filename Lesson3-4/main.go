package main

func main() {
	Init()
	for ; BreakJudge(); day++ {
		Morning()
		Noon()
		fg := Evening()
		CheckAchieve()
		if fg == true {
			break
		}
	}
	Summary()
}
