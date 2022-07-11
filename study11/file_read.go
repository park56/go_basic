// 파일 읽기(1)
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
	// 파일 읽기
	// Open : 기존 파일 열기
	// Close : 리소스 닫기
	// Read, ReadAt : 파일 읽기
	// 권한 주의, 예외처리 중요
	// 탐색 seek 중요

	// 파일 열기
	file, err := os.Open("sample.txt")
	errCheck1(err)

	// 파일 읽기 #1
	fileInfo, err := file.Stat() // 파일 사이즈 확인을 위해 정보 획득
	errCheck2(err)
	fd1 := make([]byte, fileInfo.Size()) // 슬라이스에 읽은 내용 담기 (두번 째 인수는 슬라이스 크기).
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력(1) : ", fileInfo, "\n")
	fmt.Println("파일 이름(2) : ", fileInfo.Name(), "\n")
	fmt.Println("파일 크기(3) : ", fileInfo.Size(), "\n")
	fmt.Println("파일 수정 시간(4) : ", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업(1) 완료 (%d bytes)\n\n", ct1)
	//fmt.Println(fd1)         // 타입변경 X
	fmt.Println(string(fd1)) // 타입변경 O

	fmt.Println("==========================")

	// 읽기 #2(탐색 : seek(offset))
	ol, err := file.Seek(20, 0) // 파일을 읽을 범위 지정 0: 처음 위치, 1: 현재 위치, 2: 마지막 위치
	errCheck1(err)
	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	errCheck1(err)
	fmt.Printf("읽기 작업(2) 완료 (%d bytes) (%d ret)\n\n", ct2, ol)
	fmt.Println(string(fd2), "\n")
	fmt.Println("==========================")

	// 읽기 #3
	o2, err := file.Seek(0, 0)
	errCheck1(err)
	fd3 := make([]byte, 50)
	//ct3, err := file.ReadAt(fd3, 8) // offset 위치부터(두번째 인수) 읽어 온다.
	ct3, err := file.Read(fd3)
	errCheck2(err)
	fmt.Printf("읽기 작업(3) 완료 (%d bytes) (%d ret)\n\n", ct3, o2)
	fmt.Println(string(fd3), "\n")
	fmt.Println("==========================")
}
