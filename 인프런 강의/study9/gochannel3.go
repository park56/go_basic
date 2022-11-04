// 채널 기초(3)

package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널

	// #1  (동기 : 버퍼 미사용)
	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
			time.Sleep(20 * time.Second) // sleep 주석 처리후 테스트
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}
}
