package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// 결과 정보를 담을 구조체
type SearchResult struct {
	Filename string
	LineNum  int
	Content  string
}

func searchInFile(path string, keyword string, results chan<- SearchResult, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(path)
	if err != nil {
		return // 열 수 없는 파일은 무시
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			// 검색 성공 시 채널로 결과 전송
			results <- SearchResult{
				Filename: path,
				LineNum:  lineNum,
				Content:  strings.TrimSpace(line),
			}
		}
		lineNum++
	}
}

func main() {
	root := "./"           // 현재 폴더부터 검색
	keyword := "func"      // 찾을 단어 (예: Go 함수 선언부)
	results := make(chan SearchResult)
	var wg sync.WaitGroup

	// 1. 디렉토리 탐색 및 고루틴 생성
	go func() {
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			// 파일인 경우에만 검색 고루틴 실행
			if !info.IsDir() {
				wg.Add(1)
				go searchInFile(path, keyword, results, &wg)
			}
			return nil
		})
		if err != nil {
			fmt.Println("탐색 에러:", err)
		}

		// 모든 고루틴이 끝날 때까지 기다린 후 결과 채널 닫기
		wg.Wait()
		close(results)
	}()

	// 2. 결과 출력 (채널이 닫힐 때까지 계속 읽음)
	fmt.Printf("'%s' 검색 결과:\n", keyword)
	for res := range results {
		fmt.Printf("%s [%d라인]: %s\n", res.Filename, res.LineNum, res.Content)
	}
}