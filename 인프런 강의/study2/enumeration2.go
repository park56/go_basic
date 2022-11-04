package main

import "fmt"

func main() {
	const (
		A = iota + 1
		B
		C
		D
	)
	fmt.Println(A, B, C, D)
}
