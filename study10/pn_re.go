// Panic & Recover(1)

package main

import (
	"fmt"
	"log"
)

func main() {
	// Go panic 함수
	// 사용자가 에러 발생 가능
	// panic함수는 호출 즉시, 해당 메소드를 즉시 중지시키고 defer함수를 호출하고 자기 자신을 호출한 곳으로 리턴
	// 런타임 이외에 사용자가 코드 흐름에 따라 에러를 발생 시킬 때 중료
	// 문법적 에러는 아니지만, 논리적인 코드 흐름에 따른 에러 발생 처리 가능

	fmt.Println("Start Main")

	//panic("Error occurred : user stopped") // 방법1
	log.Panic("Error occurred : user stopped") //방법2
	fmt.Println("end main")

}
