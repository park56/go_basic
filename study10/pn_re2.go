// panic & recover (2)

package main

import "fmt"

func runFunc() {
	defer func() {
		s := recover() // recover는 panic에서 넘긴 매개변수를 가지고 있음
		fmt.Println("Erorr Message : ", s)
	}()

	panic("Error occurred")    // 밑으로는 실행되지 않음
	fmt.Println("After panic") // 실행되지 않음
}

func main() {
	// Go Recover 함수
	// 에러 복구 가능
	// Panic으로 발생한 에러를 복구 후 코드 재실행(프로그램 종료X)
	// 즉, 코드 흐름 정상상태로 복구하는 기능
	// Panic에서 설정한 메시지를 받아올 수 있다.

	// #1
	runFunc()

	fmt.Println("Hello")
}
