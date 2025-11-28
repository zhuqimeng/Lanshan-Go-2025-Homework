package main

import (
	"fmt"
	"time"
)

type Task struct {
	Runnable func(workerId int)
}

func main() {
	ch := make(chan Task, 10)
	for id := range 10 {
		go func(workerId int) {
			for t := range ch {
				t.Runnable(workerId)
			}
		}(id)
	}
	for i := range 20 {
		j := i
		t1 := Task{
			Runnable: func(workerId int) {
				fmt.Printf("workerId%v：task%v做一件事情\n", workerId, j)
			},
		}
		ch <- t1
	}
	time.Sleep(1 * time.Second)
	close(ch)
}
