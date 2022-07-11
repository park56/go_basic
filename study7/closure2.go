// 함수 Closure(2)

package main

import "fmt"

func main() {
	// #1
	cnt := increaseCnt()

	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
}

func increaseCnt() func() int {
	n := 0              // 지역변수 (캡쳐됨) 사용 후 소멸되지 않고 호출 될 때 이전의 값을 유지
	return func() int { // 이 함수에 사용되기 위해 위에 캡쳐된 변수는 삭제 되지 않음
		n++
		return n
	}
}
