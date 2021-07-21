# 일급 함수
Go 언어는 1급 함수를 지원합니다. 즉, Go에서 함수는 "일급 시민"으로 취급됩니다.  
일급 함수를 지원하는 프로그래밍 언어에서는 함수를 변수에 할당하고, 해당 변수에서 함수를 호출할 수 있습니다.  

아래 코드는 먼저 sayHi 함수를 정의합니다. main() 함수에서는 func() 타입을 가진 myFunction이라는 변수를 선언합니다. func() 타입은 해당 변수가 함수를 저장할 수 있음을 의미합니다.

그 다음 sayHi 함수 자체를 myFunction에 할당합니다. 괄호는 붙이지 않습니다. 즉, 함수를 호출하는 코드인 sayHi()의 형태로 작성하지 않습니다. 다음과 같이 함수 이름만 작성합니다.

```go
myFunction = sayHi
```

이 코드는 sayHi 함수 자체를 myFunction 변수에 할당합니다. 하지만 다음 라인에서는, 다음과 같이 myFunction 변수명 뒤에 괄호를 추가합니다.

```go
myFunction()
```

이 코드가 실행되면 myFunction 변수 안에 저장된 함수가 호출됩니다.

아래는 전체 코드의 실행 모습입니다.

```go
func sayHi() {
    fmt.Pritnln("hi")
}
func main() {
    // 함수 타입의 변수를 선언 합니다.
    var myFunction func()
    // 변수에 sayHi 객체를 할당합니다.
    myFunction = sayHi
    // 변수에 저장된 함수를 호출합니다.
    myFunction()
}
```

