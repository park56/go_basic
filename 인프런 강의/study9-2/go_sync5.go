// 고루틴 동기화 기초 (5)

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 고루틴 동기화 객체
	// 동기화 상태(조건) 메소드 사용
	// Wait, notify, notifyALL : 기타 언어
	// Wait, Signal, Broadcast

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) // 비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Waiting : ", n)
			condition.Wait() // 고루틴 대기(멈춤)
			fmt.Println("Waiting End : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
		//fmt.Println("received :", <-c)
	}

	// 하나씩 깨우기
	/*for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println("Wake Go routine(Signal) : ", i)
		condition.Signal() // 모든 고루틴 생성 후 하나씩 깨움
		mutex.Unlock()
	}*/

	// 한번에 모두 깨우기
	mutex.Lock()
	fmt.Println("Wake Go routine(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)
}
