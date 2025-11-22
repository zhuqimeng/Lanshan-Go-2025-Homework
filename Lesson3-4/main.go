package main

func main() {
	Init()
	for ; ; day++ {
		Morning()
		Noon()
		Afternoon()
		fg := Evening()
		CheckAchieve()
		if fg || BreakJudge() {
			break
		}
	}
	Summary()
}
