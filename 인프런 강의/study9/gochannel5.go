// 채널 기초(5)

package main

import (
	"fmt"
)

func main() {
	// 채널
	// Close : 채널 닫기, 주의 -> 닫힌 채널에 값 전송 시 패닉(예외) 발생
	// Range : 채널안에서 값을 꺼낸다.(순회), 채널 닫아야 반복문 종료 -> 채널이 열려있고 값 전송하지 않으면 계속 대기

	// #1
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // 5회 채널에 값 전송 후 채널 닫기
	}()

	for i := range ch { // 채널에서 값을 꺼내온다.(채널이 Close될 때까지)
		fmt.Println("ex1 : ", i)
	}
}
