// 파일 쓰기(1)
package main

import (
	"fmt"
	"os"
)

// 에러 체크 방식1
func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	// 파일 쓰기
	// Create : 새 파일 작성 및 파일 열기
	// Close : 리소스 닫기
	// Write, WriteAt, writeString : 파일 쓰기
	// 각 운영체제 권한 주의(오류 메시지 확인)
	// 예외 처리 정말 중요
	// https://golang.org/pkg/os

	// #1
	file, err := os.Create("0406file.txt")
	errCheck1(err)

	// 리소스 해제
	defer file.Close()

	// #1
	s1 := []byte{74, 78, 77, 76, 76}
	n1, err := file.Write([]byte(s1)) // 문자열 -> byte 슬라이스형으로 변환 후 쓰기
	errCheck2(err)
	fmt.Printf("1. 쓰기 작업 완료 (%d bytes)\n", n1)
	file.Sync() // Write commit(Stable)

	// #2
	s2 := "Hello Golang\n File write test -1 \n"
	n2, err := file.WriteString(s2)
	errCheck2(err)
	fmt.Printf("2. 쓰기 작업 완료 (%d bytes)\n", n2)
	file.Sync()

	// #3
	s3 := "Test WriteAt -2\n"
	n3, err := file.WriteAt([]byte(s3), 70) // len(offset) 조절하면서 테스트
	errCheck2(err)
	fmt.Printf("3. 쓰기 작업 완료 (%d bytes)\n", n3)
	file.Sync()

	// #4
	n4, err := file.WriteString("Hello golang\n File Wirite Test\n")
	errCheck2(err)
	fmt.Printf("4. 쓰기 작업 완료 (%d bytes)\n", n4)

}
