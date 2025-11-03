package main

import "fmt"

const pi float64 = 3.14

func main() {
	r := 5
	var area = pi * float64(r*r)
	fmt.Println(area)
}
