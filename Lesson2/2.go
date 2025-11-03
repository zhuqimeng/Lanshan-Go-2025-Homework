package main

import "fmt"

func main() {
	var a, b int
	var op byte
	fmt.Scanf("%d %c %d", &a, &op, &b)
	task := calucate(op)
	if task != nil {
		fmt.Println(task(a, b))
	} else {
		panic("error input")
	}
}

func calucate(op byte) func(int, int) int {
	switch op {
	case '+':
		return func(x, y int) int {
			return x + y
		}
	case '-':
		return func(x, y int) int {
			return x - y
		}
	case '*':
		return func(x, y int) int {
			return x * y
		}
	case '/':
		return func(x, y int) int {
			if y == 0 {
				panic("division by zero")
			}
			return x / y
		}
	default:
		return nil
	}
}
