# 채널과 컨텍스트

### 채널 (channel)

고루틴 간 메시지를 전달하는 메시지 큐

```go
var chan string message = make(chan string,2) // 버퍼 크기 2
```

채널을 통해 역할을 구분하여 동시성의 DeadLock을 해결할 수 있다.

채널 종료 방법
```
close를 통해 완료됐음을 알림
waitGroup의 Done을 통해 작업이 완료됐음을 알리고 Return 해버림 
```
- 고루틴 릭 : 채널을 닫지 않아 데이터를 무한정 기다리며 자원을 차지한다.


### 컨텍스트(context)

고루틴에 작업을 요청 시 작업 취소나 작업 시간등을 설정할 수 있는 작업 명세서

```
context.WithCancel()
context.Background()
context.WithCancel(context.Background())
```