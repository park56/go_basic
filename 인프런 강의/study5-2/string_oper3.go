// 문자열 연산(3)
// 중요!
package main

import (
	"fmt"
	"strings"
)

func main() {
	// #1 (결합 : 일반연산)
	str := "The Go programming language is an open source project to make programmers more productive." +
		"Go is expressive, concise, clean, and efficient." +
		"Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines," +
		" while its novel type system enables flexible and modular program construction"
	str2 := "Dong Hae Mul Gwa Back Du San Yee Ma ru go dal do rock hananim yee bow ha sa uri nara man sea"

	fmt.Println("ex1 : ", str+str2)

	// #2 (결합  : Join (추천))
	strSet := []string{} // 슬라이스 선언

	strSet = append(strSet, str)
	strSet = append(strSet, str2)
	fmt.Println("ex2 : ", strings.Join(strSet, "---")) // 두 변수 사이를 ---으로 구분 하겠다 /\ "" 는 공백으로 두겠다Devastate /\ join은 string 안에 있음
}
