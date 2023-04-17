package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestFetchGoogle(t *testing.T) {
	r, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		fmt.Println("Error : ", err)
	}

	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)

	defer cancelFunc()

	r = r.WithContext(timeoutRequest)

	_, err = http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error : ", err)
	}
}
