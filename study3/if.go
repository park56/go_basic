package main

import "fmt"

func main() {
	//제어문, 조건문
	//if문 : 반드시 bool 검사 -> 0,1,(사용불가 : 자동 형변환 불가)
	//소괄호 사용 안함
	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("15이상")
	}

	if b >= 25 {
		fmt.Pritnln("25이상")
	}

	//에러 발생하는 경우
	/*
	  #1
	  if b>= 25
	  {
	  fmt.Println("괄호는 같은 높이로")
	  }
	  #2
	  if b>24
	    fmt.Println("괄호 미사용 X")
	  #3
	  if c:=1; c{
	    fmt.Println("0,1으로 불리언 타입 변경불가")
	  }
	  #4
	  if d := 20; d>10{
	    fmt.Println("d는 10보다 크다")
	  }
	  d += 20
	  fmt.Println(d,"d는 if함수 안에서 사용되고 지워지기 때문에 오류남")
	*/

}
