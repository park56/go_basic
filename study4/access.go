package main

import (
	"fmt"
	"StudyGo/study4/lib2"
)

func main() {
	// 패키지 접근 제어
	// 변수, 상수, 함수, 메서드, 구조체 등 식별자에서 첫글자가
	// 대문자 : 패키지 외부에서 접근 가능
	// 소문자 : 패키지 외부에서 접근 불가
	fmt.Println("100보다 큰 수 : ", lib2.CheckNum(101))
	fmt.Println("100보다 큰 수 : ", lib2.CheckNum2(999))

}
