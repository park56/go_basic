package main

import (
	"fmt"
)

func init() {

	// init은 main이 없으면 실행되지 않지만 import로 호출당하면 import 될 때 같이 실행됨
	fmt.Println("init method start")
}
