package main

import "fmt"

func main() {
	// 참과 거짓
	// 조건부 논리 연산자랑 주로 사용 : !,||(or),&&(and)
	// 1,0,nil을 bool로 사용할 수 없음

	// #1
	var b1 bool = true //bool 생략가능
	var b2 bool = false
	b3 := true

	fmt.Println("ex1 : ", b1)
	fmt.Println("ex2 : ", b2)
	fmt.Println("ex3 : ", b3)

	fmt.Println("ex4 : ", b3 == b3)
	fmt.Println("ex5 : ", b1 && b3)
	fmt.Println("ex6 : ", b1 || b2)
	fmt.Println("ex7 : ", !b1)

	//b4 := 1
	// 1 != true
	/*if b4 {
		fmt.Println("b4 is True?")
	}*/

}
