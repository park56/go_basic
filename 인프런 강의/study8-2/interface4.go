// 인터페이스 기본(4)

package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// 구조체 Dog 메소드 구현

func (d Dog) run() {
	fmt.Println(d.name, ": Dog running")
}

// 구조체 Dog 메소드 구현

func (d Cat) run() {
	fmt.Println(d.name, ": Cat running")
}

func act(animal interface{ run() }) { // 익명 인터페이스(타입 정의 X)
	animal.run()
}

func main() {

	// 익명 인터페이스 사용 예제(즉시 선언 후 사용)

	dog := Dog{"poll", 10}
	cat := Cat{"navi", 5}

	// 개 행동 실행
	act(dog)
	act(cat)
}
