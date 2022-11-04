// 채널 심화(2)

package main

import (
	"fmt"
)

func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 1; i < cnt; i++ {
			sum += i
		}
		tot <- sum
	}()
	return tot
}

func main() {
	// 채널을 함수의 반환값으로 사용

	// #1
	c := sum(100)
	fmt.Println("ex1 : ", <-c)
}
