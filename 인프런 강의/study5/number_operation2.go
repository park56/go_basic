package main

import (
	"fmt"
)

func main() {
	// #1
	var n1 uint8 = 125
	var n2 uint8 = 90
	fmt.Println("ex1 : ", n1+n2)
	fmt.Println("ex1 : ", n1-n2)
	fmt.Println("ex1 : ", n1*n2)
	fmt.Println("ex1 : ", n1/n2)
	fmt.Println("ex1 : ", n1%n2)
	fmt.Println("ex1 : ", n1<<2)
	fmt.Println("ex1 : ", n1>>2)
	fmt.Println("ex1 : ", ^n2) // not 보수처리할 때 사용

	// #2
	var n3 int = 12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000
	fmt.Println("ex2 : ", float32(n3)+n4) // flaot으로 바꿔서 할 때
	fmt.Println("ex2 : ", n3+int(n4))     // int로 바꿔서 할
	fmt.Println("ex2 : ", n5+uint16(n6))  // uint16에서 120000은 허용범위를 넘기 때문에 강제로 맥스값으로 변경 후 계산

}
