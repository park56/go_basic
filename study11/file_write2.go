package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

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
	// CSV 파일 쓰기 예제
	// 패키지 저장소를 통해서 Excell 등 다양한 파일 형식 쓰기, 읽기 가능
	// 패키지 Github주소 : https://github.com/tealeg/xlsx
	// bufio : 파일의 용량이 클 경우 버퍼 사용 권장

	file, err := os.Create("text_write.csv")
	errCheck1(err)

	defer file.Close()

	// csv writer 생성
	wr := csv.NewWriter(file)
	// wr := csv.NewWriter(bufio.NewWriter(file))

	// csv 내용 쓰기
	wr.Write([]string{"Kim", "4.8"})
	wr.Write([]string{"Lee", "5.8"})
	wr.Write([]string{"Jun", "1.6"})
	wr.Write([]string{"Park", "3.4"})
	wr.Write([]string{"Jin", "6.1"})

	wr.Flush() // 버퍼 -> 파일로 쓰기

	fi, err := file.Stat()
	errCheck2(err)

	fmt.Printf("CSV 쓰기 작업 후 파일 크기(%d bytes)\n", fi.Size())
	fmt.Println("CSV 파일명 : ", fi.Name())
	fmt.Println("CSV 권한 : ", fi.Mode())

}
