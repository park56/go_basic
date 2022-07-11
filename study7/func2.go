// 함수 기초(2)

package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a, b int) { // 두 매개변수의 타입이 같을 때
	fmt.Println("ex1 : ", a+b)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i = *i * 10 // *i *= 10
}

func main() {
	// 함수(콜백) 전달, 참조 전달(call by reference), 값 전달(call by value)

	// #1
	// call back
	sum(100, add)

	// #2
	// call by value
	a := 100
	multi_value(a)
	fmt.Println("ex2 : ", a)

	// #3
	// call by reference
	b := 100
	multi_reference(&b)
	fmt.Println("ex3 : ", b)
}
