// 사용자 정의 타입(4)

package main

import "fmt"

type shopBas struct{ cnt, price int }

// 결제 함수
func (b shopBas) purchase() int {
	return b.cnt * b.price
}

// 원본 수정(참조)
func (b *shopBas) repurchase(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

// 원본 수정 X (값 전달 형식)
func (b shopBas) Repurchase(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	// 리시버(구조체 메소드) 전달(값, 참조) 형식
	// 함수는 기본적으로 값 호출 -> 변수의 값 복사 후 내부 전달(원본 수정 X) -> 맵, 슬라이스 등은 참조 전달
	// 리시버(구조체)도 포인터를 사용해 원본 수정 가능

	// #1
	bs1 := shopBas{3, 5000}
	fmt.Println("ex1(totalPrice) : ", bs1.purchase())

	// 참조형 (원본 수정 됨)
	bs1.repurchase(7, 50000)
	fmt.Println("ex2(totalPrice) : ", bs1.purchase())

	// 값 전달형 (원본 수정 안됨)
	bs1.Repurchase(7, 50000)
	fmt.Println("ex3(totalPrice) : ", bs1.purchase())
}
