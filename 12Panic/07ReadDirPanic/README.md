# scanDirectory에서 panic 사용하기
코드의 scanDirectory함수는 에러값을 반환하는 대신 panic을 호출하도록 수정되었습니다. 이는 에러 처리를 아주 단순하게 만듭니다.

# How to run code
```zsh
cd $LEARNGOPATH
go run 12Panic/07ReadDirPanic/scanDir.go
```