package main

import "fmt"

func main() {
	//문장 끝 세미콜론 필요 X
	//살행파일 생성시 자동으로 세미콜론 삽입
	//두 문장을 한 문장에 표현시 세미콜론 명시적으로 사용
	//반복문, 제어문 사용시 주의
	//예제 1
	for i := 0; i <= 10; i++ {
		//fmt.Println(i," ");fmt.Println(i)
		fmt.Print("ex : ")
		fmt.Println(i)
	}
	fmt.Println()
	//예제 2
	for j := 10; j >= 0; j-- {
		fmt.Println("ex2 : ", j)
	}
}
