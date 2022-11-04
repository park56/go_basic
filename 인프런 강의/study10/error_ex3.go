// Go 에러 처리 고급(3)
package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// 예외(에러) 처리 구조체
type PowError struct {
	time    time.Time   // 에러 발생 시간
	value   interface{} // 파라미터
	message string      // 에러 메시지
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value : %g) - %s", e.time, e.value, e.message)
}

func Power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다"}
	}
	if math.IsNaN(f) {
		return 0, PowError{time: time.Now(), value: f, message: "숫자가 아닙니다"}
	}
	if math.IsNaN(i) {
		return 0, PowError{time: time.Now(), value: f, message: "숫자가 아닙니다"}
	}
	return math.Pow(f, i), nil
}

func main() {
	// Error 처리 고급
	// error 타입이 아닐때 에러처리
	// Error 메소드를 구현해서 사용자 정의 에러 처리 예제 심화
	// 구조체를 사용해 세부적인 정보 출력

	// #1
	v, err := Power(3, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ex1 : ", v)

	s, err := Power(10, 3)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err.(PowError).time)
		fmt.Println(err.(PowError).value)
		fmt.Println(err.(PowError).message)
	}
	fmt.Println("ex1 : ", s)
}
