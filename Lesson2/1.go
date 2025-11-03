package main

import "fmt"

func main() {
	var arr []int
	var n, x int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		arr = append(arr, x)
	}
	fmt.Println(count(arr))
}

func count(slice []int) map[int]int {
	nums := make(map[int]int, 100)
	for _, value := range slice {
		nums[value]++
	}
	return nums
}
