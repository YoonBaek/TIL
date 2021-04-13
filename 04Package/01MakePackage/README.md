# Understanding GO Workspace
Go Workspace should have 3 sub directories.  
Usually, Go installer set this automatically,  
but if not, execute this code in your terminal.  
```zsh
mkdir $GOPATH/bin pkg src
```
Usually, GOPATH is your $HOME/go.

# 1. bin
"Executable" compiled binary programs are saved.

# 2. pkg
Compiled binary packages are saved.

# 3. src
Go source code is located.