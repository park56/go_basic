package gateway

type RedisGateway interface {
	SetData(key string, value string) error
	GetData(key string) (string, error)
	GetKeyList() ([]string, error)
}

// IOC를 위한 인터페이스
/*IOC - 제어 반전
보통의 흐름 -> 내 코드에서 외부 라이브러리 코드를 호출
IOC -> 외부 코드에서 내 라이브러리를 호출

이유 :
작업 구현 방식과 수행을 분리
모듈을 바꿔도 다른 시스템에 부작용이 없음
다른 시스템의 동작에 고민할 필요 없이 미리 정해진 협약대로만 동작하면 됨
모듈을 제작할 때 *모듈의 목적애 집중할 수 있다.

모듈 -
프로그램의 구성 요소, 관련된 데이터를 하나로 묶은 단위
*/
