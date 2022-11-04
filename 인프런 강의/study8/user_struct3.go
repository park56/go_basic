// 사용자 정의 타입(3)

package main

import (
	"fmt"
)

type totlCost func(int, int) int

func describe(cnt int, price int, fn totlCost) {
	fmt.Printf("cnt : %d, price : %d, orderprice : %d", cnt, price, fn(cnt, price))
}

func main() {
	// 함수 사용자 정의 타입

	var orderPrice totlCost
	orderPrice = func(cnt, price int) int {
		return (cnt * price) + 100000
	}
	describe(3, 5000, orderPrice)
}
