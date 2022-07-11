package main

import (
	checkUp "StudyGo/study4/lib" //lib 대신 checkUp으로 함수를 불러올 수 있음
	_ "StudyGo/study4/lib2"      // _을 붙이면 사용하지 않아도 없어지지 않음
	"fmt"
)

func main() {
	// 패키지 접근제어
	// 별칭 사용
	// 빈 식별자 사용
	fmt.Println("10보다 큰 수?: ", checkUp.CheckNum(11))
}
