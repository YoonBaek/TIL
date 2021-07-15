# Git

버전 관리 도구

## Why?

기존 버전관리 : 파일명으로만 진행했었음

### 자료간의 차이를 명확히 알기 어렵다.

졸업논문 1차.hwp
졸업논문 2차.hwp
졸업논문 최종.hwp
졸업논문_진짜최종.hwp
...

### 자료마다 용량이 추가 됨.

동일파일이라도 버전마다 새 파일이 생성되면 리니어하게 용량 증가.
변경 사항에 대한 기록만 남기자는 아이디어

# How?

## Git 관리

### Git 시작

``` zsh
$ git init 
```

### 스냅샷

```zsh
$ git add .
# git commit -m "message"
$ git commit -m 'first commit'
```

### 원격 저장소와 연결 및 업로드

```zsh
$ git remote add origin https://github.com/USERNAME/repositoryNAME.git
$ git push origin main
```

```zsh
$ git log
$ git status
```

