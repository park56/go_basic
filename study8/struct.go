// 구조체 기본 (1)

package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	// Go 구조체 : 필드 O, 메소드 X
	// + 객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연결
	// 상속, 객체,, 클래스 개념 X
	// 구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	// #1
	kim := Account{number: "245-901", balance: 1000000, interest: 0.015}
	lee := Account{number: "245-902", balance: 1200000} // 0으로 초기화
	park := Account{number: "245-903", interest: 0.025}
	cho := Account{"245-904", 1500000, 0.03}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)
	fmt.Println("ex1 : ", park)
	fmt.Println("ex1 : ", cho)

	// #2
	fmt.Println("ex2 : ", int(kim.Calculate()))
	fmt.Println("ex2 : ", int(lee.Calculate()))
	fmt.Println("ex2 : ", int(park.Calculate()))
	fmt.Println("ex2 : ", int(cho.Calculate()))
}
