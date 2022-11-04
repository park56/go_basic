package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//예제 1
	rand.Seed(time.Now().UnixNano())  //rand의 생성기준(seed)를 현제 시간으로 주는것
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, "50이상 100미만")
	case i >= 25 && i < 50:
		fmt.Println("i->", i, "25이상 50미만")
	default:
		fmt.Println("i->", i, "기본값")
	}
}
