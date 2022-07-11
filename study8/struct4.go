// 구조체 기본(4)

package main

import "fmt"

func main() {
	// 구조체 익명 선언 및 활용

	// #1  type 구조체명 타입
	car := struct{ name, color string }{"520d", "red"}
	fmt.Println("ex1 : ", car)
	fmt.Printf("ex2 : %#v\n", car)

	// #2
	cars := []struct{ name, color string }{{"520d", "red"}, {"530i", "yellow"}, {"578k", "white"}}

	for _, c := range cars {
		fmt.Printf("(%s),(%s) _____ (%#v)\n", c.name, c.color, c)
	}
}
