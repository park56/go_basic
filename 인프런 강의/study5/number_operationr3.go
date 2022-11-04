package main

import "math"

func main() {
	// #1 오버플로우 에러 : 범위 초과
	var n1 uint8 = math.MaxUint8 + 1
	var n2 uint16 = math.MaxUint16 + 1
	var n3 uint32 = math.MaxUint32 + 1
	var n4 uint64 = math.MaxUint64 + 1

	// #2  오버플로우 에러 : 범위 미달
	var n5 uint8 = -1
	var n5 uint16 = -1
	var n5 uint32 = -1
	var n5 uint64 = -1
}
