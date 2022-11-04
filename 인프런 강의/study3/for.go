package main

import "fmt"

func main() {
	//go에서 유일한 반복문
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

  //range 용법
  loc := []string{"seoul", "Busen", "Daegu"}
  for index, name := range loc {    //range는 무조건 인덱스, 값의 순서로 반환
    fmt.Println(index, name)
  }

	//에러 발생 예제
	//무한 루프
	/*
		 for {
			 fmt.Println("It's golang")
			 fmt.Println("infinite loop")
		 }
	*/
}
