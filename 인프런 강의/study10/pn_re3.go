// panic & recover (3)

package main

import "fmt"

func runFunc() {
	defer func() {
		if s := recover(); s != nil { // panic되서 recover에 값이 담기게 되었을 때
			fmt.Println("Erorr Message : ", s)
		}
	}()

	a := [3]int{1, 2, 3}

	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", a[i]) // index 시스템적으로 자동으로 panic 처리 됨
	}
}

func main() {
	// Go Recover 함수

	// #1
	runFunc()

	fmt.Println("Hello")
}
