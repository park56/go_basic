// 채널 심화(1)

package main

import (
	"fmt"
	"time"
)

func sendOnly(c chan<- int, cnt int) {
	for i := 0; i < cnt; i++ {
		c <- i
	}

	c <- 7777

	//fmt.Println(<-c)  발신전용에서 수신해 오류발생
}

func receiveOnly(c <-chan int) {
	for i := range c {
		fmt.Println("received : ", i)
	}

	fmt.Println(<-c)
}

func main() {
	// 채널
	// 함수 등의 매개변수에 수신 및 발신 전용 채널 지정 가능
	// 전용 채널 설정 후 방향이 다를 경우 예외 발생
	// 발신 전용 channel <- 데이터형
	// 수신 전용 <- channel
	// 매개변수를 통해 전용 채널 확인 가능
	// 채널 또한 함수의 반환 값으로 사용 가능

	// #1
	c := make(chan int)

	go sendOnly(c, 10) // 발신전용
	go receiveOnly(c)  // 수신전용

	time.Sleep(2 * time.Second) // 2초간 대기
}
