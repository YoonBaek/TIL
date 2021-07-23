# CLI (command-line-interface)

## Interface

* 서로 다른 두 개의 시스템, 장치 사이에서 정보나 신호를 주고받는 경우의 접점이나 경계면.
  * 예시 : 컴퓨터 - 모니터 : HDMInterface
* 그렇다면 사람과 컴퓨터의 접점은?
  * HardWare : 키보드, 모니터 ...
  * SoftWare : Graphic User Interface, **Command Line Interface**

## GUI와 CLI의 근본적 목적

* 컴퓨터의 조작
  * 파일의 생성, 명령, 삭제 등...

## 명령어의 종류

* Unix / Linux 계열
* Windows
  * Unix / Linux 계열의 명령어를 공부해야 한다.
  * 클라우드 컴퓨팅에 사용되는 언어의 대다수가 Unix / Linux
    * 이 원격 컴퓨터를 조작하는 방법을 배우려면 unix/linux 환경의 명령어를 학습해야 한다.

## Git Bash

1. Git을 사용하기 위한 bash shell
2. Unix / Linux 계열의 명령어를 윈도우에서 사용할 수 있도록 번역해줌

## Symbols

1. $ (Prompt) : 명령어를 받아들일 준비가 되어 있다는 신호
2. / (Root dir) : 모든 파일 / 폴더 들의 최상위 폴더
3. ~ (Home dir) : 사용자(계정)에게 할당된 폴더
4. . (pwd) : 현재 일하고 있는 폴더
5. .. : 현재 위치한 폴더를 기준으로 상위폴더를 지칭
   * 계속 진행하다보면 root에 도달
6. Tab : 파일 / 폴더 이름 자동완성 기능
7. Ctrl + c : 실행중인 프로세스 취소 (cancel)
8. Ctrl + l : 터미널 창 정리하기 (clear)
9. 방향키 위 아래 : 이전에 입력한 명령어 히스토리 탐색

## Commands

* $ ls
  * 현재 위치한 폴더 내부의 파일 / 폴더 출력
* $ ls -a
  * 현재 위치한 폴더 내부의 모든 파일/폴더 출력
* $ touch <filename> 
  * 파일 생성
* $ start <filename>
  * 윈도우에서 파일 실행
  * 맥의 경우는 open
* $ rm <filename>
  * 파일 삭제
* $ mkdir <dirname>
  * 폴더 생성
* $ cd <dirname>
  * 해당 디렉토리로 이동
* $ cd ..
  * 상위 디렉토리로 이동
* $ rm -r 
  * 해당 디렉토리와 하위 구조 삭제

## Custom Commands

* zprofile (mac) 혹은 bash_profile을 활용해서 alias 만들어서 활용할 수 있다.

  ```zsh
  # zprofile에 아래 라인 추가
  alias typora="open -a typora"
  ```

  