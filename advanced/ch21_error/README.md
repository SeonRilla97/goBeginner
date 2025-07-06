# 에러 핸들링


### error 는 인터페이스다

```go
type error interface{
	Error() string
}
```
=> 어떤 타입이든 Error() 메서드를 가지고 있다면 에러로 사용가능하다.


1. 메서드가 반환하는 error 타입
2. 직업 에러를 반환하는 메소드
3. 메소드에 에러 구현
4. 기존 에러를 감싸서 새로운 에러로 도출
    => 감싸진 에러에서 원본 에러 꺼내기 사용 : error.As() / error.Is()


### Panic

프로그램 정상 진행 힘들 때 프로그램 흐름을 중지 시킨다.

프로그램을 즉시 종료 및 에러메세지 출력하고 콜 스택(함수 호출순서)를 표시
```go
// panic 정의는 다음과 같다.
func panic(interface{})
// 주로 fmt.Errof와 같이 사용
panic(fmt.Errof("This is error num:%d",num ))
```

##### 패닉 전파

1. 콜 스택의 역순으로 에러를 전파한다.
2. recover()를 통해 복구를 시도할 수있으며, 복구 시점부터 프로그램이 계속 실행된다




