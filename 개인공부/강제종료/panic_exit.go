package main

import "os"

func ItsPanic() {
	panic("panic")
}

func ItsExit() {
	// 현제 프로세스를 즉시 종료
	// 현제 실행중인 모든 고루틴과 지연 실행 함수는 실행되지 않고 프로그램이 종료된다.
	// 종료코드는 공식적으로 있지는 않고 개발자간 협의에 의해 결정된다.
	os.Exit(1)
}
