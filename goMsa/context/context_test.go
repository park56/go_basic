package gomsa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestFetchGoogle(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://google.com", nil)

	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond) // 아웃바운드 요청에서 context 취소 단계를 수동으로 수행해야 한다.
	defer cancelFunc()

	r = r.WithContext(timeoutRequest) // 매개 ctx를 r의 얕은 복사본이 된 ctx를 반환한다. 이 함수를 실행하면 1 milli second 후 에러가 발생하며 요청이 끝난다(시간 초과). Do메서드는 즉시 리턴된다.

	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error : ", err)
	}
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

type validataionHandler struct {
	next http.Handler
}

type validationContextKey string // 컨테스트가 패키지를 넘어 전달될 때 key값의 타입이 같으면, 다른 패키지의 같은 key의 값을 변경하려는 컨텍스트가 name의 키에 값을 쓰면 의도치 않게 다른패키지가 context의 값을 덮어쓰게된다.

func (h validataionHandler) serveHTTP(rw http.ResponseWriter, r *http.Request) {

	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}
	// 유효한 요청이 있을 때 이 요청에 대한 새 컨텍스트를 만든 다음, 요청의 Name 필드값을 컨텍스트에 설정한다.
	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name) // r.Context(포인터)의 복사본이 전달된다. 이 때 사용한 포인터를 참조 해제 해야한다.
	r = r.WithContext(c)                                                            // 포인터를 업데이트하기 위해 포인터가 리턴된 값(c)를 참조하도록 해야 하며 원래 가리키던 값은 참조 해제 해야한다.

	h.next.ServeHTTP(rw, r)
}
