// 데이터 타입 : 문자열(2)

package main

import "fmt"

func main() {
	// 문자열 표현
	// 문자열 : UTF-8 인코딩 (유니코드 문자 집합) -> 바이트 수 주의

	// #1
	var str1 string = "Golang" // 배열의 형태로 저장된다고 보면 됨
	var str2 string = "world"
	var str3 string = "고프로그래밍" // 한글은 하나당 3byte라 5자리는 18byte

	fmt.Println("ex : ", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5]) // 해당하는 아스키코드값이 나옴
	fmt.Println("ex : ", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Println("ex : ", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	fmt.Println("ex : " + string(str1[0]) + string(str1[1]) + string(str1[2]) + string(str1[3]) + string(str1[4]) + string(str1[5])) // 해당하는 아스키코드값이 나옴

	// #2
	fmt.Printf("ex2 : %c %c %c %c %c %c\n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Printf("ex2 : %c %c %c %c %c\n", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Printf("ex2 : %s %s %s %s %s %s\n", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5]) // str 한글은 아톰에서 깨짐

	conStr := []rune(str3)
	fmt.Printf("ex3 : %c %c %c %c %c %c\n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5]) // 한글 선언을 rune으로 하면 안 꺠짐

	// #3
	for i, char := range str1 {
		fmt.Printf("ex4 : %c(%d)\t", char, i)
	}
	fmt.Println("")

	for i, char := range str2 {
		fmt.Printf("ex4 : %c(%d)\t", char, i)
	}
	fmt.Println("")
	for i, char := range str3 { // 이렇게 하니까 왠지 모르게 안깨짐
		fmt.Printf("ex4 : %s(%d)\t", char, i)
	}
}
