// 자료형 : 맵(2)

package main

import "fmt"

func main() {
	// 맵
	// 맵 조회 및 순회(Interator)

	// #1
	map1 := map[string]string{
		"daum":   "https://daum.net",
		"naver":  "https://naver.com",
		"google": "https://google.com",
	}

	fmt.Println("ex1 : ", map1["google"])
	fmt.Println("ex1 : ", map1["daum"])
	fmt.Println()

	// #2
	for k, v := range map1 {
		fmt.Println("ex2 : ", k, v)
	}
	fmt.Println()

	for _, v := range map1 {
		fmt.Println("ex2 : ", v)
	}
}
