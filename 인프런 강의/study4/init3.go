package main

import (
	"StudyGo/study4/lib" // lib에 init함수가 있으면 바로 실행
	"fmt"
)

var num int32

func init() {
	num = 30
	fmt.Println("Main init start")
}

func main() {
	fmt.Println("10보다 큰 수? : ", lib.CheckNum(9))
}
