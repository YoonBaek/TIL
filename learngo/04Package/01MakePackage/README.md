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

# Package naming conventions
1. Only lowercase letter is recommended for package name.
2. If clear, you can use shorter word. ex) function -> func
3. Only one word is recommended. If multi-word is enevitable, use no underline or uppercase.  
ex) strconv
4. If package name and variable name make shadowing, don't use the variable name.