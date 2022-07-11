// 사용자 패키지 작성 및 문서화 예제
package main

// alias 사용(패키지 중복 또는 약자로 사용)
import (
	oper "StudyGo/study12/arithmetic"
	"fmt"
)

func main() {
	// 사용자 패키지 작성 & Go 문서화
	// GOPATH/src
	// 패키지 폴더명(디렉토리명) 명확하게 지정
	// 패키지 파일의 package 이름으로  사용 --> 길면 alias사용
	// package main을 제외하고 package 문서에 등록  --> 모두 패키지를 사용
	// 기본적으로 GOROOT (pkg)에서 검색 후 없으면 GOPATH (src/pkg)검색
	// go install 명령어 실행시 -> GOPATH/pkg에 등록(문서화)
	// godoc - http=:6060(임의의 포트) -> pkg이동 -> 본인패키지 메소드 및 주석 확인(pkg,type,method) 주석

	// 패키지 사용 #1
	nums := oper.Numbers{100, 10}
	fmt.Println("Package Used(1) : ", nums.Plus())
	fmt.Println("Package Used(2) : ", nums.Minus())
	fmt.Println("Package Used(3) : ", nums.Multi())
	fmt.Println("Package Used(4) : ", nums.Divide())
	fmt.Println("Package Used(5) : ", nums.MultiPlus())
	fmt.Println("Package Used(6) : ", nums.MultiMinus())

}
