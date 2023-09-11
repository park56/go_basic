// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("err: ", err)
// 		}
// 	}()

// 	getPanic()

// 	fmt.Println("복구됨")
// }

// func getPanic() {
// 	panic(errors.New("get panic"))
// }
