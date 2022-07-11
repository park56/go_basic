// 자료형 : 슬라이스(1)
package main

import (
	"fmt"
)

func main() {
	// 길이(가변) -> 동적으로 크기가 늘어남, 레퍼런스(참조 값) 타입
	// 슬라이스 (길이&용량) 크기가 동적으로 할당

	// 2가지 선언 방법 1. 배열처럼 선언 2. make 함수 : make(자료형, 길이, 용량(생략시 길이로))

	// 관련함수
	// cap() : 배열, 슬라이스 용량
	// len() : 배열, 슬라이스 개수

	// #1  배열처럼 선언 (초기화가 되지않음)
	var slice []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	// slice2[0] = 10    // 아직 초기화를 하지않아 변경불가
	slice3[4] = 10 // 값 수정 가능

	fmt.Printf("%-5T %d %d %v\n", slice, len(slice), cap(slice), slice)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	// #2 make로 선언  (초기화가 됨)
	var slice5 []int = make([]int, 5, 10)
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)

	slice6[2] = 7

	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

	// #3  배열처럼 선언 nil확인
	var slice9 []int // nil 슬라이스(길이와 용량 0)

	if slice9 == nil {
		fmt.Println("This is nil slice")
	}
}
