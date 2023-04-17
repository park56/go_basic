package main

import (
	"errors"
	"os"

	"example.com/m/v2/calcu"
)

func main() {

	outer := os.Stdout

	SM := calcu.SuccessMessage{
		Message: "성공적인 메세지",
	}

	BM := calcu.BadMessage{
		Message: "실패한 메세지",
		Erorr:   errors.New("실패했어용").Error(),
	}

	SM.SendSuccess(*outer)
	BM.SendBad(*outer)
}
