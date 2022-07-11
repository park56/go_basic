package main

import "fmt"

func main() {
	//_처리로 골라내며 열거하기
	const (
		_ = iota + 0.75*2
		Default
		Gold
		Platinum
	)
	fmt.Println(Default, Gold, Platinum)
}
