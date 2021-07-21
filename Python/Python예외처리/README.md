# 예외처리

## 디버깅

> "코드의 상태를 신중하게 출력해가며 심사숙고하는 것보다 효과적인 디버깅 도구는 없습니다." - 브라이언 커니핸, Unix for Beginners

* print문 활용
  * 특정 함수 결과, 반복 / 조건 결과등 나눠서 생각
* 개발환경(text editor, IDE)등에서 제공하는 기능 활용
  * breakpoint, 변수 조회 등
* Python tutor 활용 (단순 파이썬 코드인 경우)
* 뇌 컴파일, 눈 디버깅

## 코드를 작성하다가...

* 에러 메시지가 발생하는 경우
  * 해당하는 위치를 찾아 메시지를 해결
* 로직 에러가 발생하는 경우
  * 명시적인 에러 메시지 없이 예상과 다른 결과가 나온 경우
    * 정상적으로 동작하였던 코드 이후 작성된 코드를 생각해 봄
    * 전체 코드를 살펴봄
    * 휴식을 가져봄
    * 누군가에게 코드를 설명해봄

## 문법 에러(SyntaxError)

* SyntaxError가 발생하면, 파이썬 프로그램은 실행이 되지 않음
* 파일이름, 줄 번호, ^ 문자를 통해 파이썬이 코드를 읽어 나갈 때 문제가 발생한 위치를 표현
* 줄에서 에러가 발생된 가장 앞 위치를 나타낼 때 캐럿 ^ 문자 사용하여 표시
* Invalid syntax
* assign to literal
* EOL (end of line)
* EOF (end of file)
* 등이 있음. 공통점으로는 프로그램이 실행되지 않음

## 예외 (Exception)

* 실행 도중 예상치 못한 상황을 맞이하면, 프로그램 실행을 멈춤
  * 문장이나 표현식이 문법적으로 올바르더라도 발생하는 에러
* 실행 중에 감지되는 에러들을 예외(Exception)라고 부름
* 예외는 여러 타입(type)으로 나타나고, 타입이 메시지의 일부로 출력됨
  *  NameError, TypeError 등은 발생한 예외 타입의 종류
* 모든 내장 예외는 Exception Class를 상속받아 이뤄짐
* 사용자 정의 예외를 만들어 관리할 수 있음

## 예외 종류

* ZeroDivisionError : 0으로 나누고자 할 때 발생
* NameError : namespace 상에 이름이 없는 경우
* TypeError : 
  * 타입 불일치
  * argument 누락
  * argument 개수 초과
  * arguement 타입 불일치

* ValueError : 타입은 올바르나 값이 적절치 않거나 없는 경우
* IndexError : 인덱스가 존재하지 않거나 범위를 벗어나는 경우
  *  자주 보게 되는 유형의 에러
* KeyError : 해당 키가 존재하지 않는 경우
* ModuleNotFoundError : 존재하지 않는 모듈을 import 하는 경우
* ImportError : 모듈은 있으나 존재하지 않는 클래스/함수 등을 가져오는 경우
* KeyboardInterrupt : 임의로 프로그램을 종료하였을 때
* IndentationError : 공백이 적절하지 않는 경우

## 파이썬 내장 예외(built-in-exceptions)

* Python.org docs에서 확인 가능

# 예외 처리

## 예외 처리

* try문(statement) / except절 (clause)를 이용하여 예외 처리를 할 수 있음.

* try 아래의 코드블록이 실행됨

  * 예외가 발생되지 않으면, except 없이 실행종료

  * 예외 발생시, except 발생 후 실행 종료

    ```python
    try :
      	num = input("숫자 입력 :")
        print(int(num))
    except ValueError :
      	print("숫자가 아닙니다.")
    # 숫자 입력 : 안녕
    # 숫자가 아닙니다.
    ```

### 복수의 예외처리 실습

* 100을 사용자가 입력한 정수로 나누고 출력하는 코드를 작성해보시오.

  ```python
  try :
    	num = input("값을 입력하시오:")
    	100 / int(num)
  except ValueError :
      # 위에서부터 처리하기 때문에 가장 작은 범주부터 예외처리 해야한다.
  except ZeroDivisionError :
    	# code
  ```

## 예외 처리 종합 예시

* 파일을 열고 읽는 코드를 작성하는 경우

  * 파일 열기 시도
    * 파일 없는 경우 => '해당 파일이 없습니다'(except)
    * 파일 있는 경우 => 파일 내용을 출력(else)

  * 해당 읽기 작업 종료 메시지 출력(finally)

    ```python
    try :
      	f = open("noofile.text")
    except FileNotFoundError : # 예외 O
      	print("해당 파일이 없습니다.")
    else : # 예외 X
      	print("파일을 읽기 시작합니다.")
        print(f.read())
        print("파일을 모두 읽었습니다.")
        f.close()
    finally : # 무조건 실행
      print("읽기를 종료합니다.)
    ```

  ## 에러 메시지 처리

  * as 키워드를 활용하여 원본 에러 메시지를 사용할 수 있음

    * 예외를 다른 이름에 대입

      ```python
      try :
        	empty_list = []
          print(empty_list[-1])
      except IndexError as err :
        	print(f"{err} 오류가 발생했습니다.")
      # list index out of range 오류가 발생했습니다.
      ```

  # 예외 발생시키기

  ## raise

  * raise를 통해 예외를 강제로 발생

    ```python
    raise ValueError("값 에러 발생")
    ```

  ## assert

  * assert를 통해 예외를 강제로 발생

  * assert는 상태를 검증하는데 사용되며, 무조건 AssertionError가 발생

  * 일반적으로 디버깅 용도로 사용

    ```python
    assert len([1, 2]) == 1. "길이가 1이 아닙니다."
    # AssertionError 길이가 1이 아닙니다.
    ```

  * 파이썬 실행
    * -O 옵션으로 파이썬을 실행하는 경우, assert문과 \_\_debug\_\_ 에 따른 조건부 코드를 제거 후 실행

# try / except 자주 써도 될까?

## try

* EAFP

  * Easier to ask for forgiveness than permission.
  * 허락을 구하는게 용서를 받는게 더 쉽다
  * 일단 저지르는게 낫다 라는게 파이썬의 이념
  * 예외처리를 활용하여 검사를 수행하지 않고 일단 실행하고 예외처리를 진행하는 스타일
  * 파이썬 코드가 문제 없이 "실행될 것"을 전제로 코드 실행

  ```python
  try :
    	x = my_dict['key']
  except KeyError:
    	pass
  ```

## if

 * LBYL

   	*  Look Before You Leap
      	*  어떤 것을 실행하기 전에 에러가 날만한 요소들을 조건문으로 검사를 하고 수행

   ```python
   if 'key' in my_dict :
     	x = my_dict['key']
   else :
     	pass
   ```