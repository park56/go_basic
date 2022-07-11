package main

import "fmt"

func main() {
	// 문자열 연산
	// 추출, 비교, 조합(결합)

	// #1  (추출 : 슬라이스)
	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println("ex : ", str1[0:2], str2[0])
	fmt.Println("ex : ", str2[3:], str2[0])
	fmt.Println("ex : ", str2[:4])
	fmt.Println("ex : ", str1[1:3])

}
