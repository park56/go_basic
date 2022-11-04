// 문자열 연산(2)

package main

import "fmt"

func main() {
	str1 := "Golang"
	str2 := "World"

	fmt.Println("ex : ", str1 == str2)
	fmt.Println("ex2 : ", str1 != str2)
	fmt.Println("ex3 : ", str1 > str2)
	fmt.Println("ex3 : ", str1 < str2) // Go 문자열 -> 아스키 코드에 대한 사전식 비교 /\ 아스키코드에서 뒤쪽에 있으면 더 큰거라고 판단
}
