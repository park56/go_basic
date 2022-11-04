// 라이브러리 접근제어 (1)
package lib

import "fmt"

func CheckNum(c int32) bool { // CheckNum이 checkNum이 되면 못찾음
	return c > 10
}

func init() {
	fmt.Println("lib package > init start")
}
