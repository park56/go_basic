package gomsa

import (
	"context"
	"time"
)

type contextStr string

// context타입은 여러 go 루틴에서 동시에 안전하게 사용할 수 있는 요청 범위 데이터애 접근하기 위한 안전한 방법을 구현한다.

func my_Context() {
	ctx := context.Background() // 빈 context를 반환한다. 보통 main 함수에서 최상위 context로 사용된다.

	ctx, cancel := context.WithCancel(ctx) // cancel함수를 가지고 있는 부모 context의 복사본을 리턴한다. cancel함수를 호출하면 context와 연결된 리소스가 해제되기 때문에 실행중인 작업이 완료된 직후 호출되야 한다.
	cancel()

	ctx, cancel = context.WithDeadline(ctx, time.Now().Add(1*time.Microsecond)) //	제한시간이 지난 경우 만료되는 부모 context의 복사본을 리턴한다. 제한시간이 지나면 context의 완료 채널이 닫하고 연결된 리소스가 해제 된다.
	cancel()

	ctx, cancel = context.WithTimeout(ctx, 3*time.Minute) // WithDeadline과 유사하나 매개변수로 지속시간을 전달한다. 시간이 지나면 완료채널이 닫히고 리소스가 해제된다.
	cancel()

	ctx = context.WithValue(ctx, contextStr("key"), "val") // 전달받은 키와 값이 쌍을 가지고 있는 부모 context의 복사본을 리턴한다. 이렇게 설정된 context의 값은 요청 범위 데이터로 사용하기에 적합하다. 외부 패키지로 전달될 일이 많이 때문에 사용자 설정 타입을 사용하는편이 좋다.

	ctx.Done()
}
