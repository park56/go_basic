// 자료형 : 슬라이스 심화(2)
package main

import (
	"fmt"
	"sort"
)

func main() {

	// 슬라이스 추출 및 정렬
	// slice[i:j]  i -> j-1
	// slice[i:]   i -> 마지막
	// slice[:j]  처음 -> j-1
	// slice[:]   싹다

	// #1 (추출)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("ex1 : ", slice[:])
	fmt.Println("ex1 : ", slice[0:])
	fmt.Println("ex1 : ", slice[:5])
	fmt.Println("ex1 : ", slice[0:len(slice)])
	fmt.Println("ex1 : ", slice[3:])
	fmt.Println("ex1 : ", slice[:3])
	fmt.Println("ex1 : ", slice[1:3])
	fmt.Println()

	// #2 (정렬)
	// sort 패키지 : https://golang.org/pkg/sort

	slice2 := []int{3, 6, 9, 1, 2, 4, 6, 8, 7, 10}
	slice3 := []string{"b", "a", "f", "d", "e", "c"}

	fmt.Println("ex2 : ", sort.IntsAreSorted(slice2)) //정렬을 확인라는 메소드
	sort.Ints(slice2)                                 // 정렬하는 메소드
	fmt.Println("ex2 : ", slice2)
	fmt.Println("ex2 : ", sort.IntsAreSorted(slice2))
	fmt.Println()

	fmt.Println("ex3 : ", sort.StringsAreSorted(slice3))
	sort.Strings(slice3)
	fmt.Println("ex3 : ", slice3)
	fmt.Println("ex3 : ", sort.StringsAreSorted(slice3))

}
