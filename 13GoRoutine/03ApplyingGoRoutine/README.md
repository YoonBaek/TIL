# responseSize함수에 고루틴 사용하기.
```go
for _, url := range urls {
    go responseSize(url)
}
```
앞선 intro의 코드에서 go만 추가해주면 go routine의 수를 늘릴 수 있습니다!  
간단하죠?

하지만 모든 각 responseSize가 끝날 때 까지 기다리지 않으면
main 함수 고루틴이 종료되면서 같이 끝나기 때문에 기다려줘야합니다.  
main 고루틴이 종료되어버리면 프로그램도 끝나기 때문입니다.
```go
time.Sleep(time.Second)
```
그러나 기다리는 시간은 누가 정해줄 것이며, 불확실한 상태에서 무작정 오래 기다린다면, go routine의 빠르기를 장담할 수 있을까요?

**time.Sleep()**은 그렇게 좋은 방안이 아니어 보입니다.

# 고루틴 실행 순서
```
Getting https://www.daum.net...
Size : 237258
Getting https://www.naver.com...
Size : 212414
Getting https://www.google.com...
Size : 14024
```
매번 코드를 실행할 때마다 위 순서가 바뀌는 것을 보실수 있습니다. 즉, 일반적으로는 go routine은 실행 순서를 보장하지 않습니다. 하지만 실행 순서가 중요하다면, chan을 활용할 수 있습니다.
