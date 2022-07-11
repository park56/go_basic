// 채널 기초(4)

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 채널

	// #1  (비동기 : 버퍼 사용)
	// 버퍼 : 발신 -> 가득차면 대기, 비어있으면 작동, 수신 -> 비어있으면 대기, 가득 차있으면 작동
	runtime.GOMAXPROCS(4)
	ch := make(chan bool, 4)
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
			//	time.Sleep(20 * time.Second) // sleep 주석 처리후 테스트
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}
}
