// 함수 심화(1)
package main

import "fmt"

func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}
	return tot
}

func sum(a ...int) int {
	ssum := 0
	for _, i := range a {
		ssum += i
	}
	return ssum
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2 : ", value)
	}
}

func prtWord1(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex3 : ", value)
	}
}

func main() {
	// 함수 고급
	// 가변 인자 (매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않을 때)

	// #1
	x := multiply(1, 2, 3, 4)
	fmt.Println("ex1 : ", x)
	y := sum(7, 7, 7, 7, 7, 7, 7, 7, 7, 7)
	fmt.Println("ex1 : ", y)
	fmt.Println()

	// #2
	prtWord("apple", "pine", "pineApple", "strowberry", "melon", "waterMelon", "banana")
	fmt.Println()

	// #3
	a := []int{1, 2, 3, 4, 5}
	arr := []string{"Park", "Kim", "Han", "Kwon", "Ryo"}
	m := multiply(a...)
	n := sum(a...)

	fmt.Println("ex3 : ", m)
	fmt.Println("ex3 : ", n)
	prtWord1(arr...)

}
