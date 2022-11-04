package main

import "fmt"

func main() {
	// 슬라이스 (슬라이스 참조 타입 증명)

	// #1 (배열)
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int
	arr2 = arr1
	arr2[0] = 7

	fmt.Print("ex1 : ", arr1, "\nex1 : ", arr2, "\n")

	// #2 (슬라이스)
	slice1 := []int{1, 2, 3}
	var slice2 []int
	slice2 = slice1
	slice2[0] = 7
	fmt.Print("ex1 : ", slice1, "\nex1 : ", slice2, "\n")

	// #3 (슬라이스 예외 상황)
	slice3 := make([]int, 50, 100)
	fmt.Println("ex3 : ", slice3[4])
	//	fmt.Println("ex3 : ", slice3[50]) // make에서 정한 길이 만큼만 초기화가 되기 때문에 이걸 하면 out of range가 나옴

	// #4 (슬라이스 순회)
	for i, v := range slice3 {
		fmt.Println(i, v)
	}
}
