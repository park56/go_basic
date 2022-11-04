package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 데이터 타입 : 문자열(1)
	// 문자열
	// 큰 따옴표 "", 백스쿼트 ``
	// GOlang : char 타입 존재 하지 않음 -> rune(int32)로 문자 코드 값으로 표현
	// 문자 : ''로 작성
	// 자주 사용하는 escape : \\, \', \", \a(콘솔벨), \b(백스페이스), \f(쪽바꿈), \n(줄바꿈), \r(복귀), \t(탭) 등등

	// #1
	var str1 string = "c:\\go_study\\src\\"
	str2 := `c:\go_study\src\` // 백스쿼트 안쪽은 그대로 사용
	fmt.Println("ex1 : ", str1)
	fmt.Println("ex1 : ", str2)

	// #2
	var str3 string = "Hello World!"
	var str4 string = "안녕하세요."
	var str5 string = "\ud55c\uae00" // \u 유니코드 출력 (안씀)

	fmt.Println("ex2 : ", str3)
	fmt.Println("ex2 : ", str4)
	fmt.Println("ex2 : ", str5)

	// #3
	// 길이 (바이트 수)
	fmt.Println("ex3 : ", len(str3)) // 영어(1) + 특수문자(1)
	fmt.Println("ex3 : ", len(str4)) // 한글(3) + 특수문자(1)

	// #4
	// 진짜 길이
	fmt.Println("ex4 : ", utf8.RuneCountInString(str3)) // 문자열의 길이재는 함수
	fmt.Println("ex4 : ", utf8.RuneCountInString(str4))
	fmt.Println("ex4 : ", len([]rune(str4))) // rune형 배열을 선언하고 안에 str4를 넣고 길이를 잼
}
