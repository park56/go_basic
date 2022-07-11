// 자료형 : 슬라이스 심화(3)

package main

import (
	"fmt"
)

func main() {
	// 슬라이스 복사
	// copy(복사 대상, 원본)
	// make로 공간을 할당 후 복사 해야함.
	// 복사 된 슬라이스 값 변경해도 원본에는 영향 없음

	// #1 (복사)
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println("ex1 : ", slice2)
	fmt.Println("ex1 : ", slice3) // 복사 안됨(공간 할당이 이미 되어 있어야 함)
	fmt.Println()

	// #2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5) //공간할당

	copy(b, a) // 복사됨

	b[0] = 7
	b[4] = 10
	fmt.Println("ex2 : ", a)
	fmt.Println("ex2 : ", b)
	fmt.Println()

	// #3
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] // 주의! 슬라이스 추출 복사는 참조 -> 복사본 변경 시 원본값도 변경

	d[1] = 7
	fmt.Println("ex3 : ", c)
	fmt.Println("ex3 : ", d)

	// #4
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] // :7은 용량 지정

	fmt.Println("ex4 : ", len(e), cap(e), e)
	fmt.Println("ex4 : ", len(f), cap(f), f)
}
