package main

import (
	"fmt"
	"sync"
)

var (
	cnt  int
	wg   = sync.WaitGroup{}
	lock = sync.Mutex{}
	ch   = make(chan worker, 10)
)

type worker struct {
	work func()
}

type Worker interface {
	Solve(int)
	Assign(int)
}

func (w worker) Solve(num int) {
	for range num {
		wg.Add(1)
		go func() {
			for v := range ch {
				v.work()
			}
			wg.Done()
		}()
	}
}

func (w worker) Assign(num int) {
	for range num {
		tmp := worker{
			work: func() {
				lock.Lock()
				defer lock.Unlock()
				cnt++
			},
		}
		ch <- tmp
	}
	close(ch)
}

func main() {
	myWorker := worker{}
	myWorker.Solve(10)
	myWorker.Assign(50)
	wg.Wait()
	fmt.Println(cnt)
}
