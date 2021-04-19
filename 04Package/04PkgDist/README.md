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
go install $GOPATH/src/omedirectory
```

# Package Distribution
## Import pkg
You can import your own package.  
But there are a lot of packages out there which has same name as your package.
So you can import like this.
```go
// in case your package name is keyboard.
import github.com/yourName/BiggerPkg/keyboard
//...
keyboard.YourFunc()
// ...
```
## Download pkg
you can download customized package using 'go get'.  
It goes like this.
```zsh
go get github.com/YoonBaek/learngo/04Package/01MakePackage/greeting
```