// 채널 기초(6)

package main

import (
	"fmt"
)

func main() {
	// 채널
	// Close : 채널 닫기

	ch := make(chan int)

	go func() {
		for i := 0; i < 4; i++ {
			ch <- 7777
		}
	}()

	val1, ok1 := <-ch //

	fmt.Println("ex1 : ", val1, ok1)
	val2, ok2 := <-ch
	fmt.Println("ex1 : ", val2, ok2)
	val3, ok3 := <-ch
	fmt.Println("ex1 : ", val3, ok3)

	close(ch)
	val4, ok4 := <-ch
	fmt.Println("ex1 : ", val4, ok4)

}
