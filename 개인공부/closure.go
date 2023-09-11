package main

// 클로저란 리턴을 익명 함수로 할 때 상위 변수를 끌어와 데이터를 본인이 가지고 있는것처럼 사용하는 기능이다.
// 리턴된 함수안에서 캡쳐되어 사용하는 것이기 때문에 함수를 다시 실행할 경우 값이 초기화 된다.

func closure() func() int {
	// 여기의 i가 캡쳐됨
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	myClosure := closure()
	println(myClosure()) // 1
	println(myClosure()) // 2
	println(myClosure()) // 3

	myClosure = closure() // 리셋
	println(myClosure())  // 1
	println(myClosure())  // 2
	println(myClosure())  // 3
}
