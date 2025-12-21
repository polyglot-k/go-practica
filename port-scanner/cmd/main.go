package main

import (
	"fmt"
	"time"

	scanner "github.com/polyglot-k/go-practica/port-scanner/internal"
)

func main() {
	var host string
	var startPort, endPort int
	var timeoutMs int

    fmt.Print("대상 호스트: ")
    if _, err := fmt.Scanln(&host); err != nil {
        fmt.Println("잘못된 호스트 입력:", err)
        return
    }
    fmt.Print("포트 범위(시작 끝): ")
    if _, err := fmt.Scanln(&startPort, &endPort); err != nil {
        fmt.Println("잘못된 포트 범위 입력:", err)
        return
    }
    fmt.Print("타임아웃(ms): ")
    if _, err := fmt.Scanln(&timeoutMs); err != nil {
        fmt.Println("잘못된 타임아웃 입력:", err)
        return
    }

	timeout := time.Duration(timeoutMs) * time.Millisecond
	openPorts := []int{}
	startTime := time.Now()

	fmt.Printf("\n🔍 %s 스캔 시작...\n", host)

	for port := startPort; port <= endPort; port++ {
		fmt.Printf("\r스캔 중... [%d/%d]", port, endPort)

		// 분리한 패키지의 함수 호출
		if scanner.ScanPort(host, port, timeout) {
			openPorts = append(openPorts, port)
			fmt.Printf("\n[+] %d번 포트 열림\n", port)
		}
	}

	fmt.Printf("\n\n✅ 완료! 소요 시간: %v\n", time.Since(startTime))
	fmt.Printf("✅ 열린 포트 목록: %v\n", openPorts)
}