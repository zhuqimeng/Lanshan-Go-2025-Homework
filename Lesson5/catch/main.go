package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"catch/catch/pool"
	"catch/catch/search"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("使用方法: %s [目录] [关键词]\n", os.Args[0])
		os.Exit(1)
	}

	rootDir := os.Args[1]
	keyword := os.Args[2]

	// 检查目录是否存在
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		fmt.Printf("错误: 目录 '%s' 不存在\n", rootDir)
		os.Exit(1)
	}

	// 递归获取所有文件路径
	filePaths, err := getAllFiles(rootDir)
	if err != nil {
		fmt.Printf("错误: 无法读取目录: %v\n", err)
		os.Exit(1)
	}

	if len(filePaths) == 0 {
		fmt.Println("未找到任何文件")
		return
	}

	fmt.Printf("在目录 '%s' 中找到 %d 个文件，正在搜索关键词: '%s'\n\n",
		rootDir, len(filePaths), keyword)

	// 创建协程池
	workerPool := pool.NewWorkerPool(10, len(filePaths)) // 10个worker

	// 创建搜索器
	searcher := search.NewSearcher(keyword)

	var wg sync.WaitGroup
	results := make(chan *search.SearchResult, 100)

	// 启动结果收集器
	wg.Add(1)
	go func() {
		defer wg.Done()
		searcher.CollectResults(results)
	}()

	// 提交搜索任务到协程池
	for _, filePath := range filePaths {
		task := func() {
			searcher.SearchInFile(filePath, results)
		}
		workerPool.Submit(task)
	}

	// 等待所有任务完成
	workerPool.Wait()
	close(results)

	// 等待结果收集完成
	wg.Wait()

	fmt.Printf("\n搜索完成!\n")
}

// getAllFiles 递归获取目录下所有文件的路径
func getAllFiles(rootDir string) ([]string, error) {
	var files []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过目录和隐藏文件
		if !info.IsDir() && !strings.HasPrefix(info.Name(), ".") {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}
