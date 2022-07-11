// 자료형 : 포인터 (1)

package main

import "fmt"

func main() {
	// 포인터
	// 변수의 지역성, 연속된 메모리의 참조..., 힙, 스택
	// 주소의 값은 직접 변경이 불가능
	// *(애스터 리스크)으로 사용
	// nil로 초기화

	// #1
	var a *int
	var b *int = new(int)

	fmt.Println(a) // &
	fmt.Println(b)

	i := 7

	a = &i // & 주소값 전달
	b = &i

	fmt.Println("ex1 : ", a, &i) // a = i의 주소를 저장
	fmt.Println("ex1 : ", &a)    // i의 주소를 저장한 a의 주소
	fmt.Println("ex1 : ", *a)    // (*)역참조 /\저장된 주솟값으로 가서 참조

	fmt.Println()

	fmt.Println("ex1 : ", b, &i)
	fmt.Println("ex1 : ", &b)
	fmt.Println("ex1 : ", *b)

	c := &i
	d := &i

	*d = 10

	fmt.Println("ex2 : ", *c)
	fmt.Println("ex2 : ", *d)

}
