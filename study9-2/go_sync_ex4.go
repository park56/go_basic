// 고루틴 동기화 고급(4)

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// 고루틴 동기화 고급
	// 원자성 사용 -> 기능적으로 분할 불가능한 완전 보증된 일련 조작, 모두 성공하거나 모두 실패
	// 모든 조작이 완료 될 때까지 다른 프로세스 개입 불가
	// sync/atomic에서 원자적 연산자 제공
	// httpsL//golang.org/pkg/sync/atomic 공식문서
	// 주로 공용 변수에 관한 계산 사용

	// 원자성 사용 안할 경우 예제

	runtime.GOMAXPROCS(10)
	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			//cnt++
			atomic.AddInt64(&cnt, 1)
			wg.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			//cnt--
			atomic.AddInt64(&cnt, -1)
			wg.Done()
		}(i)
	}

	// Add(7000) == Done(7000) 성립해야함
	wg.Wait()

	finalCnt := atomic.LoadInt64(&cnt)
	// 방법1
	fmt.Println("WaitGroupEnd >>>> ", cnt)
	// 방법2 *추천
	fmt.Println("WaitGroupEnd >>>> ", finalCnt)
}
