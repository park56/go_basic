package main

import (
	"fmt"
	"math"
)

func main() {
	// 숫자 연산 (산술, 비교)
	// 타입이 같아야 가능 (다른것 끼리 할 시 에러발생)
	// +, -, *, %, /, >>, <<, &, ^

	// #1
	var n1 uint8 = math.MaxUint8 //Uint로 가능한 가장 큰 수
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("ex : ", n1)
	fmt.Println("ex : ", n2)
	fmt.Println("ex : ", n3)
	fmt.Println("ex : ", n4)

	fmt.Println("ex2 : ", math.MaxInt8)
	fmt.Println("ex2 : ", math.MaxInt16)
	fmt.Println("ex2 : ", math.MaxInt32)
	fmt.Println("ex2 : ", math.MaxInt64)

	fmt.Println("ex3 : ", math.MaxFloat32)
	fmt.Println("ex3 : ", math.MaxFloat64)

	n5 := 10000
	n6 := int16(10000)
	n7 := uint8(100)

	fmt.Println("ex4 : ", n5+int(n6))
	fmt.Println("ex4 : ", n6 > int16(n5))
	fmt.Println("ex4 : ", n6-int16(n7) < int16(n5))

}
