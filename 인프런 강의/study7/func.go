// 함수 기초 (1)
package main

import "fmt"
import "strconv"

// 함수 선언 위치는 상관 X
func hello() {
	fmt.Println("ex1 : Hello Wow")
}

func one_sel(m string) {
	fmt.Println("ex2 : ", m)
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	// 선언
	// func 함수이름(매개변수) 반환타입 or 반환 값 변수명 : 반환값이 존재할 때
	// func 함수이름() 반환타입 or 반환 값 변수명 : 매개변수 X, 반환 값 O
	// func 함수이름(매개변수) : 매개변수 O, 반환값 X
	// func 함수이름() : 매개변수 X, 반환값 X
	// 반환 값 다수 가능

	// #1
	hello()
	// #2
	one_sel("Hello")
	// #3
	result := sum(5, 10)
	fmt.Println("ex3 : ", result)
	fmt.Println("ex3 : ", sum(5, 10))
	fmt.Println("ex3 : ", strconv.Itoa(sum(5, 5)))
}
