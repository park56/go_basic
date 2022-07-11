// 자료형 : 포인터(3)

package main

import "fmt"

func rptc(n *int) {
	*n = 77
}

func vptc(n int) {
	n = 77
}

func main() {
	// 포인터의 값  전달
	// 함수, 메서드 호출 시 매계변수 값을 복사 전달 -> 함수, 메서드 내에서 원본 값 수정 불가
	// 매계값 변경을 위해 포인터로 전달
	// 특히  크기가 큰 배열 같은 것은 시스템에 무리 -> 포인터 전달로 해결(슬라이스, 맵은 기본이 참조라 ㄱㅊ)

	var a int = 10
	var b int = 10

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", b)

	rptc(&a)
	vptc(a)

	fmt.Println("ex2 : ", a)
	fmt.Println("ex2 : ", b)

}
