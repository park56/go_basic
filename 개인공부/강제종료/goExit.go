package main

import (
	"fmt"
	"runtime"
)

func GOExit() {
	// runtime.Goexit(): 현제 고루틴을 종료시키고 다른 실행 가능한 고루틴이 없을 때 프로그램을 종료시킨다.
	// 다른 고루틴에는 영향을 주지 않는다.
	go func() {
		go func() {
			fmt.Println("Before runtime.Goexit")
			runtime.Goexit()                    // 현재 고루틴 종료
			fmt.Println("After runtime.Goexit") // 실행되지 않음
		}()

	}()

	fmt.Println("main continue")
}
