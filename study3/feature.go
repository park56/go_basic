package main

import "fmt"

func main() {
	// GO : 모호하거나 애매한 문법이 거의 없음
	// 후위 연산자는 존재하지만 전위 연산자는 존재하지 않음
	// 증감연산은 반환값이 존재하지 X
	// 포인터를 사용하지만 연산은 사용하지 않음

	sum, i := 0, 0

	for i <= 100 {
		//sum += i++  //예외 발생
		sum += i
		i++
		//++i  //예외 발생
	}
	fmt.Println(sum)
}
