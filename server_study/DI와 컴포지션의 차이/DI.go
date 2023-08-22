package main

import "fmt"

// Greeter 인터페이스 정의
type Greeter interface {
	Greet() string
}

// EnglishGreeter 구조체 정의
type EnglishGreeter struct{}

func (eg EnglishGreeter) Greet() string {
	return "Hello"
}

// KoreanGreeter 구조체 정의
type KoreanGreeter struct{}

func (kg KoreanGreeter) Greet() string {
	return "안녕하세요"
}

// Greet 함수는 Greeter 인터페이스에 의존성을 주입받음
func Greet(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	englishGreeter := EnglishGreeter{}
	koreanGreeter := KoreanGreeter{}

	Greet(englishGreeter)
	Greet(koreanGreeter)
}
