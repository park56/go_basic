// 함수 심화(3)

package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("ex2 : (", n, ")", "hi")
	prtHello(n - 1)
}

func main() {
	// 재귀 함수 (Recursion)
	// 장점 : 프로그램이 보기 쉽고, 코드 간결, 오류 수정이 용이
	// 단점 : 코드 이해 어려움, 역량에 따라 기억공간 많이 사용, 무한루프 가능성

	// #1  (재귀 함수)
	fac := fact(4)
	fmt.Println("ex1 : ", fac)

	// #2
	prtHello(10)
}
