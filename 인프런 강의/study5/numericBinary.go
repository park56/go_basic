package main

import (
	"fmt"
	"strconv"
)

func main() {

	/* 10 -> 2 */
	var exam10_1 int64

	exam10_1 = 10
	fmt.Println("10진수 ", exam10_1)
	exam2_1 := strconv.FormatInt(exam10_1, 2)
	fmt.Println("2진수", exam2_1)

	/* 2 -> 10 */
	var exam2_2 string

	exam2_2 = "1010"
	fmt.Println("2진수 " + exam2_2)
	exam10_2, err := strconv.ParseInt(exam2_2, 2, 64)
	fmt.Println("10진수", exam10_2)
	if err != nil {
		panic(err)
	}
}
