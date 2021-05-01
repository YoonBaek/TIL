# Go install

## What's different from go build?
go install saves bin file at go workspace 'bin' folder.  
but go build saves at current directory.  

It's easier to manage bin file.

## Difference 2
go install looks always go workspace src directory.  
And it directory name should be passed.
``` zsh
# case go build
go build somegocode.go
# case go install
go install $GOPATH/src/somedirectory
```

# Package Distribution
## Import private pkg
You can import your own package.  
But there are a lot of packages out there which has same name as your package.
So you can import like this.
```go
// in case your package name is keyboard.
import github.com/yourName/BiggerPkg/keyboard
//...
keyboard.YourFunc(...)
// ...
```
## Download private pkg
you can download customized package using 'go get'.  
It goes like this.
```zsh
go get github.com/YoonBaek/learngo/04Package/01MakePackage/greeting
```
## Document your package
go 소스코드에 주석 처리하기
```go
// Package documentation
package yourpkg

// YourOwnFunc function documentation
func YourOwnFunc() {

}
```
### How to see document
쉘에서 go doc 명령어 활용
```zsh
# Package documentation
go doc github.com/yourfolder/yourpkg
# package function documentation
go doc github.com/yourfolder/yourpkg YourOwnFunc
```
# Documentation Rules
1. Must be a perfect sentence
2. Must start with "Package" and package name
3. Function documentation must start with func name
4. Using tab, you can include example code.