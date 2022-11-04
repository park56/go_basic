package main

import "fmt"

func main() {
	//제어문 조건문 switch
	// switch 뒤 표현식 생략 가능
	//case 뒤 표현식 생략 가능
	//자동 break 때문에 fallthrouth 존재
	//type 분기 -> 값이 아닌 변수 TYpe으로 분기 가능

	//예제 1
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}
	//예제 2
	// 이런 방식을 더 선호함
	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}
	//예제 3
	switch c := "go"; c { // ; c를 하는 이유는 이뒤에서 부터는 c에 대한 조건이라는 것을 표시하기 위함 만약 없다면 계속 c == "go" 같이 해야
	case "go":
		fmt.Println("Go")
	case "java":
		fmt.Println("Java")
	default:
		fmt.Println("일치 하는 값 없음")
	}
	//예제 4
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("golang")
	case "javalang":
		fmt.Println("JavaScripts")
	default:
		fmt.Println("Language")
	}
	//예제 5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i == j:
		fmt.Println("i와 j는 같다")
	case i > j:
		fmt.Println("i가 j보다 크다")
	}
}
