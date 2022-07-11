package main

import "fmt"

func main() {
	//레이블을 이용한 for
	//예제 1
loop1: //레이블 밑에 레이블과 상관없는 것이 오면 5
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break loop1 // loop1을 기준으로 탈출
			}
			fmt.Println("ex 1 : ", i, j)
		}
	}
	fmt.Println("레이블 1 탈출")
	// 예제 2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // 건너띄기
		}
		fmt.Println("ex2 : ", i)
	}

loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue loop2 //loop2로 *사실상 continue
			}
			fmt.Println("ex3 : ", i, j)
		}
	}

}
