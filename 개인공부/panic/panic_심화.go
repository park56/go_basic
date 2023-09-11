package main

import (
	"errors"
	"fmt"
)

func main() {

	// 제대로 100까지 출력하는 코드
	for i := 0; i <= 100; i++ {
		func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("Recovered from panic:", err)
				}
			}()

			if i == 30 {
				panic(errors.New("Panic at 30!"))
			}
			fmt.Println(i)
		}()
	}

	// panic을 복구하지만 그 후로는 코드가 멈추는 코드
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("Recovered from panic:", err)
	// 	}
	// }()

	// for i := 0; i <= 100; i++ {
	// 	if i == 30 {
	// 		panic(errors.New("Panic at 30!"))
	// 	}
	// 	fmt.Println(i)
	// }
}
