// 함수 심화(4)

package main

import "fmt"

func main() {
	// 함수 고급
	// 익명 함수(Anonymous Function)
	// 즉시 실행(선언과 동시에)

	// #1
	func() {
		fmt.Println("ex1 : Anonymous Function")
	}()

	// #2    매개변수 존재
	msg := "Hello Golang"
	func(m string) {
		fmt.Println("ex2 : ", m)
	}(msg)

	// #3    매개변수 존재
	func(x, y int) {
		fmt.Println("ex3 : ", x+y)
	}(10, 5)

	// #4    매개변수, 리턴값 존재
	r := func(x, y int) int {
		return x * y
	}(10, 100)
	fmt.Println("ex4 : ", r)
}
