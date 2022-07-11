// 고루틴 동기화 기초(3)

package main

import (
	"fmt"
	"time"
)

func main() {
	// 뮤텍스 : 상호 배제 -> Thread(고루틴)들이 단독으로 실행되게 하는 기술

	// #1 동기화 사용하지 않은 경우
	// 쓰기 읽기 동작 순서가 일정하지 않아 잘못된 오류를 반환 할 가능성 증가
	data := 0

	go func() {
		for i := 1; i <= 10; i++ {
			data += 1
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read1 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(5 * time.Second)

}
