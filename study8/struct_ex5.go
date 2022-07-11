// 구조체 심화(5)
// 오버라이딩
package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

type Executives struct {
	Employee     // is a 관계  /\ 상속과 비슷함
	specialBonus float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

func (e Executives) Calculate() float64 { // 오버라이딩 : 일부만 특수한 과정을 거친 계산이 필요할 때 같은 이름의 함수를 선언해 프로그램이 자동으로 우선순위를 매겨 계산해 주는 것
	return e.salary + e.bonus + e.specialBonus
}

func main() {
	// 구조체 임베디드 메소드 오버라이딩 패턴

	// #1

	// 직원
	ep1 := Employee{"kim", 10000, 10000}
	ep2 := Employee{"park", 20000, 20000}

	// 임원
	ex := Executives{
		Employee{"lee", 30000, 30000},
		30000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))

	// Employee 구조체를 통해서 메소드 호출
	fmt.Println("ex2 : ", int(ex.Calculate()))
	// fmt.Println("ex2 : ", int(ex.Calculate() + ex.specialBonus)) // 이미 오버라이딩 된 값에 추가로 계산됨 -> 잘못된 계산
	// fmt.Println("ex2 : ", int(ex.Employee.Calculate()+ex.specialBonus))  // 쓸데 없이 길게 하면 (오버라이딩 X)
}
