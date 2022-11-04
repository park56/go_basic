package main

import "fmt"

func main() {
	// #예제 1
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += 1
	}
	fmt.Println("ex : ", sum)
	//#예제 2
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
		//sum2 += j
		//j := i++ -> 에러   go에서 후치연산은 반환값이 없어 혼자 써야함
	}
	fmt.Println("ex2 : ", sum2)
	//#예제 3  while 형태와 유사하게
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3 : ", sum3)
	//예제 4  잘 사용하지 않음
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4 : ", i, j)
	}
}
