package main

import "fmt"

func main() {
	//const 상수
	//선언 후 값 변경 불가
	//고정값 관리용
	const a string = "test1"
	const b = "test2"
	const c int32 = 10 * 10
	const f = false
	/*
	   예외가 발생할 때
	   const v = 9
	   v = 10
	*/
	fmt.Println("a = ", a, "b = ", b, "c = ", c, "f = ", f)
}
