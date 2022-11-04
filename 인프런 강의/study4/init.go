package main

import (
	"fmt"
)

func init() {
	// init : 패키지 로드시에 가장 먼저 호출
	// 가장 먼저 초기화 되는 작업 작성 시 유용
	fmt.Println("init method start")
}

func main() {
	fmt.Println("Main method Start")
}
