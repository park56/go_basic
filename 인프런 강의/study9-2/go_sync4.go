// 고루틴 동기화 기초(4)

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// RWMUTEX : 쓰기 lock -> 쓰기 시도중 다른 곳에서 이전 값을 읽으면 X, 읽기 lock, 쓰기 lock 전부 방어
	// RMUTEX : 읽기 lock -> 읽기 시도중 값 변경 방지, 쓰기 lock 방어
	runtime.GOMAXPROCS(runtime.NumCPU())
	data := 0
	mutex := new(sync.RWMutex) // var mutex = new(sync.RWMUTEX)

	go func() {
		for i := 1; i <= 10; i++ {
			// 쓰기 뮤텍스 잠금
			mutex.Lock()
			data += 1
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Millisecond)
			// 쓰기 뮤텍스 잠금해제
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read1 : ", data)
			time.Sleep(1 * time.Second)
			// 읽기 뮤텍스 잠금해제
			mutex.RUnlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
			// 읽기 뮤텍스 잠금해제
			mutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)

}
