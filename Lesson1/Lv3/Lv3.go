package main

import "fmt"

func Factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(Factorial(n))
}
