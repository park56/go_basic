package main

import "fmt"

func main() {
	const a, b int = 10, 20
	const c, d, e = true, 0.32, "tst"
	const (
		x, y int16 = 50, 90
		i, k       = "sting", 9391
	)
	fmt.Println("a = ", a, "b = ", b, "c = ", c, "d = ", d, "e = ", e, "x = ", x, "y = ", y, "i = ", i, "k = ", k)
}
