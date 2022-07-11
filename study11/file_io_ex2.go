// 파일 I/O (2)
package main

import (
	"bufio"
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 파일 읽기, 버퍼사용 쓰기 -> bufio 패키지 활용
	// ioutil, bufio 등은 io.Reader, io.Writter 인터페이스를 구현 함 -> 즉, writer, read 메소드를 사용 가능
	/*
	   type Reader interface {
	     Read(p []byte) (n int, err error)
	   }

	   type Writer interface{
	     Write(p []byte) (n int, err error)
	   }

	   즉 bufio의 NewReader, NewWriter를 통해 객체 반환 후 메소드 사용

	   bufio(Bufferd io) 패키지
	   https://golang.org/pkg/bufio

	   파일열기
	   두번째 매개변수 확인
	   https://golang.org/pkg/os/#pkg-contants
	*/

	/*
	   버퍼 상태
	   a -----> a
	   b -----> ab
	   c -----> abc
	   d -----> abcd
	   e -----> e        -----> abcd
	   f -----> ef       -----> abcd
	   g -----> efg      -----> abcd
	   h -----> efgh     -----> abcd
	   i -----> i        -----> abcdefgh
	*/

	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errCheck(err)

	// bufio 파일쓰기 #1
	wt := bufio.NewWriter(file) // Writer 반환(버퍼 사용)
	wt.WriteString("Hello Golang\nFile Write Test\n")
	wt.Write([]byte("Hello Golang\nFile Write Test2\n"))

	// 버퍼 상태
	fmt.Printf("used buffer size (%d bytes)\n", wt.Buffered())          // 내가 사용한
	fmt.Printf("Remnant used buffer size (%d bytes)\n", wt.Available()) // 남은(사용가능한)
	fmt.Printf("All used buffer size (%d bytes)\n", wt.Size())          // 전체

	wt.Flush() // 버퍼의 내용을 비우고 반영

	fmt.Println("쓰기 작업 종료.")
	fmt.Println("========================================")

	rt := bufio.NewReader(file) // Reader 반환
	fi, err := file.Stat()
	errCheck(err)

	b := make([]byte, fi.Size())

	fmt.Println("파일 정보 출력 : ", fi)
	fmt.Println("파일 이름 출력 : ", fi.Name())
	fmt.Println("파일 크기 출력 : ", fi.Size())
	fmt.Println("파일 수정 시간 : ", fi.ModTime())
	fmt.Println("=========================================")

	file.Seek(0, os.SEEK_SET)
	data, _ := rt.Read(b) // 읽기(ReadLine, ReadByte, ReadByte 등)
	// rt.Read(b)

	fmt.Printf("전체 Buffer Size : (%d bytes)\n", rt.Size())
	fmt.Printf("읽기 작업 완료 : (%d bytes)\n", data)
	fmt.Println("=========================================")
	fmt.Println(string(b))
	fmt.Println("=========================================")

	defer file.Close()
}
