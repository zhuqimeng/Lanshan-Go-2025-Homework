package main

func main() {
	Init()
	for ; ; day++ {
		Morning()
		Noon()
		fg := Evening()
		CheckAchieve()
		if fg || BreakJudge() {
			break
		}
	}
	Summary()
}
