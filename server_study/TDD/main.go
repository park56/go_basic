package main

import (
	"fmt"
)

var opmap map[string]func(int, int) int

// 첫번째 결과
// func Calcurate(oper string, a, b int) int {
// 	if oper == "+" {
// 		return a + b
// 	} else if oper == "-" {
// 		return a - b
// 	} else if oper == "*" {
// 		return a * b
// 	}
// 	return a / b
// }

// 리펙토링 (성공강화) 후 첫번째 결과
// func Calcurate(oper string, a, b int) int {
// 	rst := 0
// 	switch oper {
// 	case "+":
// 		rst = a + b
// 	case "-":
// 		rst = a - b
// 	case "*":
// 		rst = a * b
// 	case "/":
// 		rst = a / b
// 	}
// 	return rst
// }

// 리펙토링 두번째 결과 - 완성형
func initOpMap() {
	opmap = make(map[string]func(int, int) int)
	opmap["+"] = add
	opmap["-"] = sub
	opmap["*"] = mul
	opmap["/"] = div
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func Calcurate(oper string, a, b int) int {
	if v, ok := opmap[oper]; ok {
		return v(a, b)
	}
	return 0
}

func main() {
	initOpMap()
	Test()
}

func Test() {

	bol := testCalcurate("Test1", "+", 4, 9, 13)
	if !bol {
		fmt.Println("fail")
		return
	}

	bol = testCalcurate("Test2", "+", 4, 9, 13)
	if !bol {
		fmt.Println("fail")
		return
	}

	bol = testCalcurate("Test3", "-", 4, 9, -5)
	if !bol {
		fmt.Println("fail")
		return
	}

	bol = testCalcurate("Test4", "*", 4, 9, 36)
	if !bol {
		fmt.Println("fail")
		return
	}

	bol = testCalcurate("Test5", "/", 9, 4, 2)
	if !bol {
		fmt.Println("fail")
		return
	}

	fmt.Println("success")
}

func testCalcurate(testcase string, op string, a, b, exp int) bool {
	o := Calcurate(op, a, b)
	if o != exp {
		fmt.Printf("%s faild %d -> %d\n", testcase, exp, o)
		return false
	}
	return true
}
