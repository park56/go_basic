package main

import (
	"errors"
	"fmt"
)

func main() {

	// import가 필요없이 그냥 사용가능한 함수는 총 4개가 있다.

	// 포맷팅이 없는 fmt.Println()
	println()
	// 포맷팅이 없는 fmt.Print()
	print()

	// 프로그램을 패닉시킨다.
	panic(errors.New("panic"))

	// 패닉을 복구시킨다. ,, panic의 상위코드에서 구현되어야 한다.
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
}
