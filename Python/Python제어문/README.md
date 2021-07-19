# Python 제어문

1. 조건문
2. 반복문

* 파이썬은 기본적으로 위에서부터 아래로 순차적으로 명령을 수행
* 특정 상황에 따라 코드를 선택적으로 실행(분기 / 조건)하거나
  계속하여 실행(반복)하는 제어가 필요함
* 제어문은 순서도(flow chart)로 표현 가능

# 조건문

* if문은 참/거짓을 판단할 수 있는 조건식과 함께 사용

  * expression에는 참/거짓에 대한 조건식이 들어감

  * 조건이 참인 경우 이후 들여쓰기 되어 있는 코드 블록을 실행

  * 이외의 경우 else 이후 들여쓰기 되어 있는 코드 블록을 실행

    * else의 경우 선택적으로 활용 가능함.

    ``` python
    a = 5
    if a > 5 :
      	print("5 초과")
    else :
     		print("5 이하")
    print(a)
    # 1 tab = 4 spaces
    # tab은 편의를 위해 vsc에서 자동화한 기능이므로
    # 파이썬의 레벨 단위는 4 spaces로 알고 있으면 된다.
    ```

## 조건문 실습문제

* 1. practice_01.py
     * 조건문을 통해 변수 num의 값의 홀수 / 짝수 여부 출력
       * 이때 num은 input을 통해 사용자로부터 입력을 받으시오.

## 복수 조건문 / 중첩 조건문

* 복수의 조건식을 활용할 경우 elif를 활용하여 표현함.

* 또한 조건문은 조건문 내부에 중첩해서 사용 가능

  ```python
  if expression :
    	# code block
  elif expression :
    	# 복수 조건문 code block
  else :
    	# code block
      if expression :
        	# 중첩 조건문 code block
  ```



## 중첩 조건문 실습 문제

* 2. Practice_01.py

  * 아래의 코드에서 **중첩조건문을 활용하여** 미세먼지 농도(dust 값)가 300이 넘는 경우 '실외활동을 자제하세요'를 추가적으로 출력하고 음수인 경우 '값이 잘못 되었습니다'를 출력하시오.

## 조건 표현식  

* 일반적으로 조건에 따라 값을 정할 때 사용

* 삼항 연산자(Ternary Operator)로 부르기도 함

  * <true인 경우 값> if <expression> else : <false>인 경우의 값

    ``` python
    value = num if num >= 0 else -num
    ```

    절댓값을 저장하기 위한 코드.

  * 홀짝 실습

    ```python
    result = "홀수입니다." if num % 2 else "짝수입니다."
    ```

# 반복문 (Loop statement)

* While 문
  * 종료 조건에 해당하는 코드를 통해 반복문을 종료시켜야 함
* for 문
  * 반복 가능(iterable)한 객체를 모두 순회하면 종료
  * 별도의 종료 조건이 필요 없음
* 반복 제어
  * break, continue, for-else

## While

* 조건식이 참인 경우 반복적으로 코드를 실행
  = 조건이 False가 될 때까지 반복

  * 조건이 참인 경우 들여쓰기 되어 있는 코드블록이 실행됨

  * 코드 블록이 모두 실행되면 다시 조건식을 검사

  * 무한 루프를 하지 않도록 반드시 종료 조건 필요

    ```python
    a = 0
    while a < 5 :
      	print(a)
        a += 1
    print("끝")
    # 0
    # 1
    # 2
    # 3
    # 4
    # 끝
    ```

### While문 실습문제

* Practice_01.py
  * 1부터 사용자가 정한 숫자까지의 총합을 구하는 프로그램

## For문

* for문은 시퀀스(string, tuple, list, range)를 포함한 iterable한 객체를 모두 순회함

* 처음부터 끝까지 순회하므로 별도의 종료조건 필요 없음

  ```python
  for fruit in ['apple', 'mango', 'banana'] :
    	print(fruit)
  # apple
  # mango
  # banana
  ```

### for문 실습문제

* Practice_01.py
  * 사용자가 입력한 문자를 한 글자씩 출력하시오,

## 리스트 순회하기 - range

* 리스트를 순회하면서 index값을 같이 출력하기

  ```python
  members = ['민수', '영희', '철수']
  for i in range(len(members)) : 
    	print(f"{i}번 : {members[i]}")
  ```

## 리스트 순회하기 - enumerate

* 위 방법을 엘레강스하게 표현하는 법

  * Start 매개변수로 시작 숫자를 지정해줄 수 있음.

  ```python
  members = ['민수', '영희', '철수']
  for i, name in enumerate(members, start = 1) : 
    	print(f"{i}번 : {name}")
  ```

# 반복문 제어

* break
  * 반복문을 종료

* continue
  * Continue 이후의 코드 블록은 수행하지 않고, 다음 반복을 수행

* for-else
  * 끝까지 반복문을 실행한 이후에 else문을 실행
    * Break를 통해 중간에 종료되는 경우 else 문은 실행되지 않음

## break

* while에서의 사용

	```python
  n = 0
  while True :
      if n == 3 :
          break
      print(n)
      n += 1
  # 0 1 2
	```

* for 문에서도 마찬가지로 사용 가능

## continue

* continue 이후의 코드블록은 수행하지 않고, 다음 반복을 수행

  ``` python
  for i in range(6) :
    	if i % 2 == 0:
      		continue
      print(i)
      
  # 1 3 5
  ```

## else

* 끝까지 반복문을 실행한 이후에 else문을 실행.

* 즉 for 문이 끝까지 돌아야지만 실행되는 구문임.

  ```python
  apple = "apple"
  banana = "banana"
  for i in apple :
    	if i == "b" :
    	    print("b!")
        	break
  else :
    	print("No B!")
      
  for i in banana :
    	if i == "b" :
        	print("b!")
        	break
  else : 
    	print("No B!")
  # No B!
  # b!
  ```

## pass

* 아무것도 하지 않음

  * 특별히 할 일이 없을 때 자리를 채우는 용도로 사용

  * 반복문이 아니어도 사용 가능

  ``` python
  for i in range(6) :
    	if i % 2 == 0:
      		continue
      print(i)
      
  # 1 3 5
  
  for i in range(6) :
    	if i % 2 == 0:
      		pass
      print(i)
  # 1, 2, 3, 4, 5
  
  ```
