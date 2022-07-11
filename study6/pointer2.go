// 지료형 : 포인터(2)

package main

import "fmt"

func main() {

	// #1
	i := 7
	p := &i
	fmt.Println("ex1 : ", i, *p)
	*p++
	fmt.Println("ex1 : ", i, *p)
	i = 99
	fmt.Println("ex1 : ", i, *p)

}
