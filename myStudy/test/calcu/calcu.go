package calcu

import (
	"fmt"
	"os"
)

type Response interface {
	SuccessMessage
	BadMessage
}

type SuccessMessage struct {
	Message string
}

type BadMessage struct {
	Message string
	Erorr   string
}

func (s *SuccessMessage) SendSuccess(set os.File) {
	fmt.Fprintln(&set, s.Message)
}

func (b *BadMessage) SendBad(set os.File) {
	fmt.Fprintln(&set, b.Message)
}
