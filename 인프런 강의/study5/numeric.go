//데이터타입 : numeric(i)

package main

import "fmt"

func main() {
	// 데이터 타입 : 숫자
	// 정수, 실수, 복소수
	// 32bit, 64bit, unsigned(양수)
	// 정수 : 8진수(0), 16진수(0x), 10진수

	var num1 int = 17
	var num2 int = -68
	var num3 int = 0631       // 0을 붙이면 8진수
	var num4 int = 0x32fa2c75 // 0x를 붙이면 16진수

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)

}
