// 기본 라이브러리를 사용해 테스트 케이스를 작성하는 방법
package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func FetchGoogle(t *testing.T) error {
	r, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		fmt.Println("Error : ", err)
		return err
	}

	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)

	defer cancelFunc()

	r = r.WithContext(timeoutRequest)

	_, err = http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error : ", err)
		return err
	}
	return nil
}
