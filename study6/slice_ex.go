// 자료형 : 슬라이스 심화(1)

package main

import "fmt"

func main() {
	// 슬라이스 추가 및 병합
	// #1
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)      // s2에 s1슬라이스를 넣음 (...은 슬라이스 일때 필수)
	s3 = append(s2, s3[0:3]...) // s3을 추출 후 병합

	fmt.Println("ex1 : ", s1)
	fmt.Println("ex1 : ", s2)
	fmt.Println("ex1 : ", s3)

	// #2
	s4 := make([]int, 0, 5)

	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("ex2 -> len : %d, cap : %d, value : %v\n", len(s4), cap(s4), s4) // 길이가 용량을 초과하면 용량이 X 2 로 늘어남
	}
}
