// 자료형 : 맵 (3)

package main

import "fmt"

func main() {
	// 맵
	// 맵 조회시 주의할 점

	// #1
	map1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value, ok := map1["lemon"]
	value2 := map1["kiwi"]     // 키값이 없는것일 경우 기본값 리턴  (int : 0, float : 0.0, string : "")
	value3, ko := map1["kiwi"] // 두 번째 리턴은 키값의 존재 유무

	fmt.Println("ex1 : ", value, ok)
	fmt.Println("ex1 : ", value2)
	fmt.Println("ex1 : ", value3, ko)

	// #2
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist")
	}
	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : banana is not exist")
	}

	if _, ok := map1["kiwi"]; !ok { // key값의 존재 유무만 파악 (else 안쓰려고)
		fmt.Println("ex2 : kiwi is not exist")
	}
}
