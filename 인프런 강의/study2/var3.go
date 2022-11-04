package main

import "fmt"

func main() {
	//짧은 선언
	// 함수 안에서만 사용가능 -> 전역불가
	//선언하고 다시 할당시 예외발생 ex. a:=10 .... a:=20 -> 예외
	//주로 제한된 범위 내에서 사용, 사용시 이건 여기안에서만 사용한다는것을 의미
	shtvar1 := 1
	shtvar2 := 2
	shtvar3 := 3
	fmt.Println("shtvar = ", shtvar1, "shtvar2 = ", shtvar2, "shtvar3 = ", shtvar3)
	//example
	if i := 1; i < 11 {
		fmt.Println("short veri is test good")
	} else {
		fmt.Println("no its not ture")
	}
}
