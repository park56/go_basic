// 고루틴 동기화 기초(1)

package main

import (
	"fmt"
	"runtime"
)

// 공유 데이터 (구조체)
type count struct {
	num int
}

func (c *count) increment() {
	c.num += 1
}

func (c *count) result() {
	fmt.Println("c : ", c.num)
}

func main() {
	// 고루틴 동기화 예제
	// 실행 흐름 제어 및 변수 동기화 가능
	// ***공유 데이터 보호가능

	// #1 (동기화 X)
	// 시스템 전체 CPU사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0} // 구조체 초기화

	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // CPU 양보 /\ 없어도 됨
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	c.result()
}
