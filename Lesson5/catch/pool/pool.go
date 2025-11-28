package pool

import "sync"

type WorkerPool struct {
	workers   int
	taskQueue chan func()
	wg        sync.WaitGroup
}

// NewWorkerPool 创建新的协程池
func NewWorkerPool(workers, taskCount int) *WorkerPool {
	pool := &WorkerPool{
		workers:   workers,
		taskQueue: make(chan func(), taskCount),
	}

	// 启动worker
	for i := 0; i < workers; i++ {
		pool.wg.Add(1)
		go pool.worker()
	}

	return pool
}

// worker 工作协程
func (p *WorkerPool) worker() {
	defer p.wg.Done()

	for task := range p.taskQueue {
		task()
	}
}

// Submit 提交任务到协程池
func (p *WorkerPool) Submit(task func()) {
	p.taskQueue <- task
}

// Wait 等待所有任务完成
func (p *WorkerPool) Wait() {
	close(p.taskQueue)
	p.wg.Wait()
}
