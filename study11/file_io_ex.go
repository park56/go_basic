// 파일 I/O (1)
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 파일 읽기, 쓰기 -> ioutil 패키지 활용
	// 더욱 편리하고 직관적으로 파일을 읽고 쓰기 가능
	// WriteFile(), ReadFile(), ReadALL() 등 사용 가능
	// httpsL//golang.org/pkg/io/ioutil/

	s := "Hello Golang\n File Write Test\n"

	// 파일모드(chmod, chown, chgrp) -> 퍼미션(권한,허락,허가)
	// 읽기 : 4, 쓰기 : 2, 실행 : 1
	// 순서 : 소유자 그룹 기타사용자  -> 모두허가 : 777
	// https://golang.org/pkg/os/#FileMode

	err := ioutil.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
	errCheck(err)

	data, err := ioutil.ReadFile("sample.txt")
	errCheck(err)

	fmt.Println("===============")
	fmt.Println(string(data))

}