# Git

버전 관리 도구

## Why?

기존 버전관리 : 파일명으로만 진행했었음

### 자료간의 차이를 명확히 알기 어렵다.

졸업논문 1차.hwp
졸업논문 2차.hwp
졸업논문 최종.hwp
졸업논문\_진짜최종.hwp
졸업논문\_찐리얼최종.hwp
...

### 자료마다 용량이 추가 됨.

위 부분을 고쳐서 초안 - 슬라이드 교체 - 오탈자 수정 등 제목을 달아준다고 해도,
동일파일이라도 버전마다 새 파일이 생성되면 리니어하게 용량 증가.

## What

## Git

변경 사항에 대한 기록만 남기자는 아이디어

## Repositoty(저장소)

* 내부의 모든 파일 / 폴더를 관리

  ``` zsh
  $ git init
  ```

  위 명령어를 통해 디렉토리를 rep으로 업그레이드.

* WARNING
  1. Home 디렉토리에서 초기화 하지 않는다.
  2. repo 안에 repo를 만들지 않는다.
     * git init 한 공간 하위에서 다시 init하지 않는다.

### 작업공간 (working directory)

* 코드 작성 및 수정
* 분장소. 코드를 분장시키는 곳. 대부분의 시간을 여기서 보내게 됨

### 스테이지 (staging area)

* 기록될 파일들의 변경사항들을 스테이지에 올리기

### 저장소

* 스테이지 위의 변경사항들을 저장 (commit)
* 스테이지 위에서 사진을 찍는 것

## Repo 백업 공간

* Github / Gitlab /Bitbucket 크게 세 가지의 서비스가 있음.

### repo 백업

1. Local repo 생성.
2. Remote repo 생성.
3. Local repo와 Remote repo를 연결
4. Commit

## How

### 커밋 남기기 (commit)

* 모든 Commit에는 작성자의 서명이 필요

* commit은 의미 있는 작업인 한에서 세부적으로 남기는게 리팩토링에도 좋다.

  * 도장을 만들어 둡시다.

    ```zsh
    # 서명에 사용할 이름(닉네임) 설정
    $ git config --global user.name <name>
    # 이메일 설정
    $ git config --global user.email <email>
    ```
    
  
*  스테이지에 올리기, 사진찍기

    ```zsh
    # git add 현재 위치
    # 스테이지에 올리기
    $ git add .
    # 저장소에 저장하기.
    # git commit -m "message"
    $ git commit -m 'first commit'
    ```

* 작업공간과 스테이지를 모니터링할 수 있는 명령어

    ```zsh
    $ git status
    ```

* 저장소 내부의 커밋들을 확인하는 명령어

  ```zsh
  $ git log

### $ git remote add : 원격 저장소와 연결 및 업로드

```zsh
$ git remote add origin https://github.com/USERNAME/repositoryNAME.git
$ git push origin <branch> # main or master or else...
```

```zsh
$ git log
$ git status
```

### $ git pull : 변경사항 다운로드

```zsh
# 해당 repo와 연결된 인가 받은 사람들만 업로드, 다운로드 가능
$ git pull origin <branch>
# 혹은 다른 repo를 최초 1회 복사
$ git clone <URL>
```

### .git ignore : 변경사항 추적 그만 두기

* .gitignore 파일 활용

  ```
  $ code .gitignore
  
  # gitignore 파일 내에서 해당 파일의 위치를 입력
  # misc
  .DS_store
  ```

### $ git diff : 로컬에서 스테이지와의 변동사항 확인하기.

* git diff 실행시 아래와 같이 변동사항이 +-로 출력되어 확인할 수 있음.

  ```sh
  diff --git a/git/0715git_basic.md b/git/0715git_basic.md
  index 8897ce0..0201b6f 100644
  --- a/git/0715git_basic.md
  +++ b/git/0715git_basic.md
  @@ -104,7 +104,7 @@
     ```zsh
     $ git log
  
  -### 원격 저장소와 연결 및 업로드
  +### $ git remote add : 원격 저장소와 연결 및 업로드
  
   ```zsh
   $ git remote add origin https://github.com/USERNAME/repositoryNAME.git
  @@ -116,7 +116,7 @@ $ git log
   $ git status
   ```
  
  -### 변경사항 다운로드
  +### $ git pull : 변경사항 다운로드
  
   ```zsh
   # 해당 repo와 연결된 인가 받은 사람들만 업로드, 다운로드 가능
  @@ -125,3 +125,20 @@ $ git pull origin <branch>
   $ git clone <URL>
   ```
  
  +### .git ignore : 변경사항 추적 그만 두기
  +
  +* .gitignore 파일 활용
  +
  +  ```
  +  $ code .gitignore
  +
  +  # gitignore 파일 내에서 해당 파일의 위치를 입력
  +  # misc
  +  .DS_store
  +  ```
  +
  +### $ git diff : 로컬에서 스테이지와의 변동사항 확인하기.
  +
  +* git diff
  ```

  
