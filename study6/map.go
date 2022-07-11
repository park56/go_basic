// 자료형 : 맵(1)

package main

import "fmt"

func main() {
	// 맵(map)
	// 맵 : 해시테이블, 딕셔너리, Key-Value 로 자료 저장
	// 레퍼런스 타입(참조 값 전달)
	// 비교 연산자 사용 불가능 (참조 타입이라서)
	// 맵 특 : Key로 참조타입 사용 불가능,Value로는 모든 타입 사용가능
	// make 함수 및 축약(리터럴)으로 초기화 가능
	// 순서 없음

	// #1
	var map1 map[string]int = make(map[string]int) // 정석
	var map2 = make(map[string]int)
	map3 := make(map[string]int)
	fmt.Println("ex1 : ", map1)
	fmt.Println("ex1 : ", map2)
	fmt.Println("ex1 : ", map3)
	fmt.Println()

	// #2
	map4 := map[string]int{} // Json 형태
	map4["apple"] = 10       // apple은 Key, 25는 Value
	map4["banana"] = 20
	map4["orange"] = 30

	map5 := map[string]int{
		"apple":  40,
		"banana": 50,
		"orange": 60,
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 70
	map6["banana"] = 80
	map6["orange"] = 93

	fmt.Println("ex2 : ", map4)
	fmt.Println("ex2 : ", map5)
	fmt.Println("ex2 : ", map6)
	fmt.Println("ex2 : ", map6["apple"])
	fmt.Println("ex2 : ", map6["banana"])
	fmt.Println()

}
