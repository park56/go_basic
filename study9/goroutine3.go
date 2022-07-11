// 고루틴 기초(3)
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func exe(name int) {
	r := rand.Intn(100)
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>", r, i)
	}
	fmt.Println(name, " end : ", time.Now())
}

func main() {
	// 고루틴
	// 멀티 코어(다중 CPU) 최대한 활용

	runtime.GOMAXPROCS(runtime.NumCPU())                        // 현 시스템의 cpu 코어 개수 반환 후 설정
	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(0)) // 설정 값 출력

	// #1
	fmt.Println("Main Routine Start : ", time.Now())
	for i := 0; i < 100; i++ {
		go exe(i) // 고루틴 100개 생성
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Main Routine End : ", time.Now())
}
