// 고루틴 기초(1)

package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("exe1 func start ", time.Now())
	time.Sleep(1 * time.Second) // 1초
	fmt.Println("exe1 func end ", time.Now())
}

func exe2() {
	fmt.Println("exe2 func start ", time.Now())
	time.Sleep(1 * time.Second) // 1초
	fmt.Println("exe2 func end ", time.Now())
}

func exe3() {
	fmt.Println("exe3 func start ", time.Now())
	time.Sleep(1 * time.Second) // 1초
	fmt.Println("exe3 func end ", time.Now())
}

func main() {
	// 고루틴(Go routine)
	// 타 언어의 쓰레드(Thread)와 비슷한  기능
	// 생성 방법 매우 간단, 리소스 매우 적게 사용
	// -> 수많은 고루틴 동시 생성 실행 가능
	// 비동기적 함수 루틴 실행(매우 적은 용량 차지) -> 채널을 통한 통신 가능
	// 공유메모리 사용 시에 정확한 동기화 코딩 필요
	// 싱글루틴에 비해 항상 빠른 처리 결과는 아니다.

	// 멀티 쓰레드 장점과 단점
	// 장 : 응답성 향상, 자원공유를 효율적으로 활용 및 사용, 작업이 분리되어 코드 간결
	// 단 : 구현하기 어려움, 테스트 및 디버깅 어려움, 전체 프로세스의 사이드이펙트(부작용), 성능저하, 동기화 코딩 필수 숙지, 데드락(교착상태)
	// 파이썬 스레드 한번 해보면 좋음

	exe1() // 일반적으로 가장 먼저 실행
	fmt.Println("Main routine Start", time.Now())
	go exe2()
	go exe3()
	time.Sleep(4 * time.Second)
	fmt.Println("Main routine End")
}
