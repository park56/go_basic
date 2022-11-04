// 힘수 Defer(4)

package main

import (
	"fmt"
)

func start(t string) string {
	fmt.Println("start : ", t)
	return t
}

func end(t string) {
	fmt.Println("end : ", t)
}

func a() {
	defer end(start("b")) // defer는 직접 해당하는 함수만 지연시킴 /\ 중첩함수 주의
	fmt.Println("in a")
}

func main() {
	// #1
	a()
}
