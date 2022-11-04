// 자료형 : 배열(2)
package main

import "fmt"

func main() {
	// 배열 순회

	// #1
	arr := [5]int{1, 10, 100, 1000, 10000}

	// Len 길이 반복
	for i := 0; i < len(arr); i++ {
		fmt.Println("ex1 : ", arr[i])
	}

	// #2   /\ 이게 더 많이 사용됨
	arr2 := [5]int{7, 77, 777, 7777, 77777}
	for _, v := range arr2 {
		fmt.Println("ex1 : ", v)
	}
}
