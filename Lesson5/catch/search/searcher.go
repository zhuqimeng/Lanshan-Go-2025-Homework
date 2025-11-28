package search

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
)

type SearchResult struct {
	FilePath string
	LineNum  int
	Content  string
	Keyword  string
}

type Searcher struct {
	keyword    string
	matchCount int32
}

// NewSearcher 创建新的搜索器
func NewSearcher(keyword string) *Searcher {
	return &Searcher{
		keyword: keyword,
	}
}

// SearchInFile 在文件中搜索关键词
func (s *Searcher) SearchInFile(filePath string, results chan<- *SearchResult) {
	file, err := os.Open(filePath)
	if err != nil {
		// 忽略无法打开的文件（如权限问题）
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		if strings.Contains(line, s.keyword) {
			// 找到匹配，发送结果
			result := &SearchResult{
				FilePath: filePath,
				LineNum:  lineNum,
				Content:  strings.TrimSpace(line),
				Keyword:  s.keyword,
			}
			results <- result
			atomic.AddInt32(&s.matchCount, 1)
		}
	}
}

// CollectResults 收集并输出搜索结果
func (s *Searcher) CollectResults(results <-chan *SearchResult) {
	for result := range results {
		s.printResult(result)
	}
}

// printResult 格式化输出搜索结果
func (s *Searcher) printResult(result *SearchResult) {
	// 获取相对路径，使输出更简洁
	relPath, err := filepath.Rel(".", result.FilePath)
	if err != nil {
		relPath = result.FilePath
	}

	fmt.Printf("📁 %s:%d\n", relPath, result.LineNum)

	// 高亮显示关键词
	highlighted := strings.ReplaceAll(result.Content, result.Keyword,
		fmt.Sprintf("\033[1;31m%s\033[0m", result.Keyword))
	fmt.Printf("   %s\n\n", highlighted)
}

// GetMatchCount 获取匹配总数
func (s *Searcher) GetMatchCount() int32 {
	return atomic.LoadInt32(&s.matchCount)
}
