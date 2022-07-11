// 고 에러 처리 기초(3)

package main

import (
	"errors"
	"fmt"
)

func main() {
	// 에러 처리
	// Errors 패키지와 New 메소드를 활용한 에러 생성
	// Errorf보다 더 많이 사용

	var err1 error = errors.New("Error occurred -1")
	err2 := errors.New("Error occurred -2")

	fmt.Println("error 1 : ", err1)
	fmt.Println("error 1 : ", err1.Error())

	fmt.Println("error 1 : ", err2)
	fmt.Println("error 1 : ", err2.Error())
}
