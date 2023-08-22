package main

import "fmt"

// Engine 인터페이스 정의
type Engine interface {
	Start()
}

// Car 구조체 정의 (컴포지션)
type Car struct {
	Engine // Engine 인터페이스를 컴포지션
}

// ElectricEngine 구조체 정의 (Engine 인터페이스 구현)
type ElectricEngine struct{}

func (e ElectricEngine) Start() {
	fmt.Println("Electric engine started")
}

func main() {
	electricEngine := ElectricEngine{}

	car := Car{Engine: electricEngine}
	car.Start() // Car 내의 Engine을 통해 ElectricEngine의 Start 메서드 호출
}
