package main

import "fmt"

func main() {
	//fallthrouth 사용

	switch a := "go"; a {
	case "java":
		fmt.Println("java")
		fallthrough
	case "go": // 한번 뚫리면 계속 True판정
		fmt.Println("go")
		fallthrough
	case "python":
		fmt.Println("python")
		fallthrough
	case "ruby":
		fmt.Println("ruby")
	}
}
