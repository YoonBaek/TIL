# Python 기초

1. 코드 스타일 가이드
2. 변수
3. 데이터 타입
4. 타입 변환
5. 연산자
6. 표현식 / 문장
7. 컨테이너

# Python 코드 스타일 가이드

* 코드를 작성하는 컨벤션에 대한 가이드 라인
* 유명한 게 크게 두가지 있음
  * PEP8 - 파이썬에서 제안하는 스타일 가이드
  * Google Style Guide - 등 기업, 오픈 소스 등에서 사용되는 스타일 가이드.

```python
print('hello')
print("world") # 두 가지의 스트링 컨벤션 사용 -> 불편

if True :
  print(True)
else :
  	print(False) # 인덴테이션을 지키지 않는 문제.
```

## 주석

* 한줄 주석은 #으로 표현
* 여러줄의 주석은 한 줄 씩 # 을 사용하거나, """, '''로 표현.
* 다만 여러 줄 주석 방법은 주로 doc을 만들 때 사용.

### Docstring

* 함수 / 클래스의 설명서, 가이드 라인 등을 작성

## 코드 라인

* 코드는 1줄에 1문장이 원칙

* 문장은 파이썬이 실행 가능한 최소한의 코드 단위

  * 기본적으로는 세미콜론(;)을 작성하지는 않음

  * 한 줄로 표기하길 원할 때 세미콜론을 작성하여 표기는 가능

  * 그러나 굳이 이렇게 자주 하지는 않는다.

# 변수와 식별자

## 변수

* 변수는 할당 연산자(=)를 통해 값을 할당(assign)
* type()
  * 변수에 할당된 값의 타입 확인

* Id()
  * 변수에 할당된 값의 고유 아이덴티티 값. 메모리 주소를 출력해서 확인 가능

* 같은 값을 동시에 할당할 수 있음

  ```python
  x = y = 10004
  ```

* 다른 값도 동시에 할당할 수 있음(multiple assignment)

  ```python
  x, y = 1, 2
  ```

  

## 값 Swap

* x, y의 값을 각각 바꿔서 저장

  ```python
  x, y = 10, 20
  
  tmp = x
  x = y
  y = tmp
  print(x, y)
  # 20, 10
  ```

* Pythonic한 방법

  ```python
  y, x = x, y
  ```

## 식별자

* 변수의 이름을 어떻게 지을까?
* 변수, 함수, 모듈, 클래스 등을 식별하는데 사용하는 이름(name)

### 규칙

* 식별자의 이름은 영문 알파벳, 언더바(_), 숫자로 구성

* 첫 글자에 숫자가 올 수 없음

* 길이 제한이 없고, 대소문자를 구별

* 다음의 키워드(keywords는) 예약어(reserved words)로 사용할 수 없음

* 확인할 키워드 모듈을 제공

  ```python
  import keyword
  print(keyword.kwlist)
  ```

  ```python
  ['False', 'None', 'True', 'and', 'as', 'assert', 'async', 'await', 'break', 'class', 'continue', 'def', 'del', 'elif', 'else', 'except', 'finally', 'for', 'from', 'global', 'if', 'import', 'in', 'is', 'lambda', 'nonlocal', 'not', 'or', 'pass', 'raise', 'return', 'try', 'while', 'with', 'yield']
  ```

* 내장 함수나 모듈 등의 이름을 지정하면 안 됨
  * 기존의 기능을 사용할 수 없게 됨.

# 데이터 타입

## 숫자(Number)

### int

* 모든 정수의 타입은 long 타입이 없고 모두 int로 표기
  * 여타 프로그래밍 언어나 Python2 에서의 long은 OS기준 32 / 64 비트

* 매우 큰 수를 나타낼 때 오버플로가 발생하지 않음

  * Overflow : 데이터 타입 별로 사용할 수 있는 메모리의 크기를 넘어서는 상황

  * Arbitrary Precision Arithmetic(임의 정밀도 산술)을 통해
    고정된 형태의 메모리가 아닌 가용 메모리들을 활용하여 모든 수 표현에 활용.

    ```Python
    import sys
    print(sys.maxsize)
    # 9223372036854775807
    # 64 비트 상에서 사용할 수 있는 최대 숫자값
    
    print(sys.maxsize ** 3)
    # 784637716923335095224261902710254454442933591094742482943
    
    print(type(sys.maxsize ** 3))
    # <class 'int'>
    ```

### float

* 정수가 아닌 실수는 float 타입

* 부동소수점

  * 실수를 컴퓨터가 표현하는 방법 - 2진수(비트)로 숫자를 표현

  * 이 과정에서 floating point rounding error가 발생하여, 예상치 못한 결과가 발생

    ```python
    # 왼쪽과 오른 쪽은 같은 값일까요?
    print(3.14 - 3.02 == 0.12 )
    ```

    ```Python
    False
    ```

    ```python
    print(3.14 - 3.02)
    ```

    ```python
    0.1200000000000001
    ```

    무한히 내려가는 실수를 완벽히 계산할 수 없기 때문에 최대한 근사한 비트로 계산하기 때문.

  * 따라서 실수는 동등 연산자로 판별할 수 없다.

  * 판별하는 방법

    ```python
    # python 3.5 이상
    import math
    math.isclose(3.14 - 3.02, 0.12)
    ```

    ```python
    True
    ```

### complex

* 실수부와 허수부로 구성된 복소수는 모두 complex타입

  * 허수부를 j로 표현

    ```python
    a = 3 + 4j
    print(type(a))
    # <class 'complex'>
    print(a.real)
    # 3.0
    print(a.imag)
    # 4.0
    ```

    

## 문자열(String)

* 모든 문자는 str타입
* 문자열은 작은 따옴표(')나 큰 따옴표(")를 활용하여 표기
  * 규정이 정해진 것은 아니나 하나만 사용하도록 PEP8에서 권장

### Escape Sequence

* 문자열 내에서 특정 문자나 조작을 위해서 역슬래시(\)을 활용하여 구분

  | 예약문자 |    내용     |
  | :------: | :---------: |
  |    \n    |   줄 바꿈   |
  |    \t    |     탭      |
  |    \r    | 캐리지 리턴 |
  |    \0    |  널 (Null)  |
  |    \'    |   싱글 퀕   |
  |    \"    |  더블 퀕\|  |

  ```python
  print('이것은 \'인용구입니다.\'')
  # 이것은 '인용구입니다.'
  ```

### String Interpolation

* 변수의 값을 문자열의 자리 표시자로(placeholder)로 대체하는 방법(과정)

  * % formatting

    ```python
    print("Hello, %s" % name)
    ```

  * str.format()

    ```python
    print("Hello, {}, Your Grade is {}").format(name, score)
    ```

  * f- strings (python 3.6+)

    ```python
    print(f'Hello, {name}, Your Grade is {score}')
    ```

* f-string 사용 예

  ```python
  import datetime
  today = datetime.datetime.now()
  print(today)
  # 2021-07-19 10:31:17.793312
  print(f"오늘은 {today:%y}년 {today:%m}월 {today:%d}일")
  # 오늘은 21년 07월 19일
  ```

## Boolean

* True / False 값을 가진 타입은 bool
* 비교 / 논리 연산에서 사용됨
* 다음은 모두 False 취급
  * 0, 0.0, (), [], {}, '', None

## None

* 값이 없음을 표현하기 위한 타입인 NoneType

  ```python
  print(type(None))
  # <class 'NoneType'>
  ```

# 타입 변환

## 자료형 변환 / 타입 변환 (Type conversion, Typecasting)

* 파이썬에서 데이터 타입은 서로 변환할 수 있음

  * 암시적 타입 변환(Implicit)

    * 사용자가 의도하지 않고, 파이썬 내부적으로 타입 변환하는 경우

      ```python
      True + 3
      ```

      4

      ```python
      3 + 5.0
      ```

      8.0

      ```python
      3 + 4j + 5
      ```

      (8 + 4j)

  * 명시적 타입 변환(Explicit)
    * 사용자가 특정 함수를 활용하여 의도적으로 타입 변환 하는 경우

## 명시적 타입 변환

작성된 형식이 맞아야 변환 가능

* int

  * str, float

  * 단, 형식에 맞는 문자열만 정수로 변환 가능

    ```python
    int('3.5') + 5
    # ValueError: invalid literal for int() with base 10: '3.5'
    ```

* float
  * str, int
  * 역시 형식에 맞아야 함

# 연산자

## 산술 연산자

| 연산자 |   내용   |
| :----: | :------: |
|   +    |   덧셈   |
|   -    |   뺄셈   |
|   *    |   곱셈   |
|   /    |  나눗셈  |
|   //   |    몫    |
|   **   | 거듭제곱 |

* 나눗셈 기본 결과는 항상 실수이나 몫 연산의 경우 정수 출력

  ```python
  print(5/2)
  print(5//2)
  print(divmod(5, 2)) # 몫과 나머지 동시 출력
  ```

  2.5

  2

  (2, 1)

## 비교 연산자

* 값을 비교하며 True / False 값을 리턴함

| 연산자 |   내용   |
| :----: | :------: |
|   <    |   미만   |
|   <=   |   이하   |
|   >    |   초과   |
|   >=   |   이상   |
|   ==   |   같음   |
|   !=   | 같지않음 |
|   is   |   객체   |
| is not |  ~객체   |

* none

  ```python
  # 특정 변수가 비어있는지 확인하기 위해서는 pep8에서는 is None을 사용하도록 권장
  a = 3
  print(a is None)
  # False
  ```

## 논리 연산자

| 연산자  |              내용              |
| :-----: | :----------------------------: |
| A and B |   A와 B 모두 True일 때, True   |
| A or B  | A와 B 모두 False 일 때, False  |
|   Not   | True를 False로, False를 True로 |

```python
print(not 0)
# True
print(not "YoonBaek")
# False
```

* 일반적으로 비교연산자와 함께 사용됨

  ``` python
  num = 3
  print(num > 3 and num < 100)
  # False
  ```

### 단축 평가

* and 연산은 좌변이 True일 때 우변을 확인하고 우변 값을 반환
  * False이면 우변을 확인하지 않고 좌변을 그냥 반환
* or 연산은 좌변이 True일 때 우변을 확인하지 않고 좌변을 그냥 반환
  * False이면 우변까지 확인하고 우변을 반환
* 굳이 불필요한 연산을 줄이기 위한 연산의 결과

```python
print(5 and 4)
print(5 or 3)
print(0 and 5)
print(5 or 0)
# 4, 5, 0, 5
```

## 복합 연산자

| 연산자  |    내용    |
| :-----: | :--------: |
| a += b  | a = a + b  |
| a -= b  | a = a - b  |
| a *= b  | a = a * b  |
| a /= b  | a = a / b  |
| a //= b | a = a // b |
| a %= b  | a = a % b  |
| a **= b | a = a ** b |

* 복합 연산자는 연산과 대입이 함꼐 이뤄짐

  * 예시) 반복문을 통해서 개수를 카운트하는 경우

    ```python
    cnt = 0
    while cnt < 3 :
      print(cnt)
      cnt += 1
    # 0
    # 1
    # 2
    ```

## Concat

* +는 숫자가 아닌 자료형에서도 사용 가능함

  ```python
  print("hello, "+"YoonBaek!")
  # hello, YoonBaek!
  ```

## Containment Test

* 포함 여부 확인

  ```python
  'a' in 'apple' # True
  ```

## Identity

* is 연산자를 활용해 동일한 객체(object)인지 확인 가능함. 

  * 자주 쓰는 -5 ~ 256 까지는 동일 id를 부여

  ```python
  a, b = 3, 3
  print(a is b)       # True
  print(id(a), id(b)) # True
  c, d = 257, 257
  print(a is b)       # True
  print(id(a), id(b)) # False
  ```

## Indexing / Slicing

* []을 통해 접근하고. [:]를 통해 슬라이싱 가능함.

  ```python
  text = "YoonBaek"
  print(text[0]) 	 # Y
  print(text[1:5]) # oonB
  ```

## 연산자 우선순위

1. ()
2. Slicing
3. Indexing
4. **
5. 단항연산자(+,-) : 부호
6. 산술연산자(*, /, %, //)
7. 산술연산자(+, -)
8. 비교연산자, ㅑㅜ, ㅑㄴ
9. not
10. and
11. or

# 표현식 / 문장

## 표현식

* 표현식은 평가(evaluate)되고, 값으로 변경

* 하나의 값으로 환원(reduce)될 수 있는 문장이어야 함

* 식별자, 값, 연산자로 구성

  ```python
  "yoonbaek" # 이것도 문장이 될 수 있습니다. 문장이 더 큰 범주.
  a = 1 + 5 # 할당문은 표현식이 아닌 문장
  a > 5 # True로 환원되는 표현식. 표현식은 문장에 속합니다.s
  ```

# 컨테이너

## 컨테이너(Container)

* 여러 개의 값을 저장할 수 있는 것 (객체)
* 시퀀스(Sequence)형 : 순서가 있는(ordered) 데이터
  * 순서가 있다 != 정렬되어 있다.
  * list, tuple, range, string, binary
* 비시퀀스(non-sequence)형 : 순서가 없는(unordered) 데이터
  * set, dictionary

## 시퀀스형 컨테이너

### list

* 순서가 있는 시퀀스로 인덱스를 통해 접근

* 순서는 0부터 시작

* 대괄호([])혹은 list()를 통해 생성

* 값에 대한 접근은 list[i]

  ```python
  a = ['a', 0, 1] # 서로 다른 타입의 원소도 저장 가능
  print(a[0]) 	  # 'a'
  
  a = [[1,2],[3,4]]
  print(a[0][1]) # 2
  ```

### tuple

* 튜플은 수정 불가능한 (immutable) 부분에서 리스트와 차이가 있다.

* 소괄호 혹은 tuple()을 통해 생성

* 값에 대한 접근은 리스트와 동일하게 my_tuple[i]

  ```python
  a = (1,2,3,1)
  a[0] = '3' # TypeError 발생. assign을 지원하지 않습니다.
  ```

* 튜플은 일반적으로 파이썬 내부에서 많이 활용됨

  * multiple assignment

  * 함수에서 복수의 값을 반환하는 경우에도 사용.

    ```	python
    x, y = 1, 2 # 여기서 1, 2는 사실 (1, 2)임.
    
    print(type(divmod(5,2)))
    # <class 'tuple'>
    ```

  * 불가변이기 때문에 내부적으로 주로 사용

* 값이 하나인 튜플 생성

  ```Python
  a = (1); b = (1,)
  print(type(a)) # int
  print(type(b)) # tuple
  ```

### range

* range는 숫자의 시퀀스를 나타내기 위해 사용

  * 기본형 : range(n)

    * 0 부터 n-1 까지의 숫자 시퀀스

  * 범위 지정 : range(n, m)

    * n부터 m-1까지의 숫자 시퀀스

  * 범위 및 스텝 지정 : range(n, m, s) :

    * n부터 m-1 까지 s만큼 증가시키면서 반환

      ```python
      print(list(range(0,10,2)))
      # [0, 2, 6, 8]
      ```

    * 역순도 가능

      ```python
      print(list(range(6,0,-1)))
      # [6, 5, 4, 3, 2, 1]
      ```

### Containment test

* 모든 시퀀스 컨테이너에서 사용 가능.

### Sequence Concat

* range 를 제외한 모든 시퀀스 컨테이너에서 사용 가능

```python
[1, 2] + ['a']
# [1, 2, 'a']
```

### Sequence Repeat

* 역시 range는 type error

```python
(1, 2) * 3
# (1, 2, 1, 2, 1, 2)
```

### Seqeunce Indexing

* 시퀀스의 특정 인덱스 값에 접근

  * 해당 인덱스가 없는 경우 IndexError

    ```python
    a = [1, 2, 3]
    a[4]
    # Index out of range
    ```

### Sequence Slicing

* range도 가능

* 역시 리스트, 튜플, 문자열 모두 가능

  ```python
  range(10)[5, 8]
  # range(5, 8)
  list(range(10))[1::3]
  # [1,4,7]
  ```

### len

* 시퀀스의 길이

* range도 적용 가능

  ```python
  len(range(10))
  # 10
  ```

### Sequence min / max

* 시퀀스의 최소, 최대 값

* 문자열의 경우 ASCII 코드 값을 통해 계산

  ```python
  max("abcdefg1234")
  # "g"
  ```

### Sequence Count

* 시퀀스 내 특정 원소의 갯수를 세줌

	```python
	text = "yoonbaek"
	text.count("o") # 2
	```

## 비시퀀스형 컨테이너

### Set

*  키 / 밸류 구분이 없기 때문에 딕셔너리와 구분 가능

* 순서가 없는 자료구조

  * 중괄호({}) 혹은 set()을 통해 생성
    * 빈 세트를 만들기 위해서는 set()을 반드시 활용해야 함
    * 중괄호만 입력하면 dict가 생성

  * 순서가 없어 별도의 값에 접근할 수 없음

* 수학에서의 집합과 동일한 구조를 가짐
  * 집합 연산이 가능
  * 중복된 값이 존재하지 않음

	```python
	blank = {}
	type(blank) # dict
	blank = set()
	type(blank) # set
	s = {1,2,3,1,2}
	print(s, type(s)) # {1,2,3} set
	```

* 집합 연산자를 사용할 수 있음.

  ```python
  set_a = {1,2,3}
  set_b = {3,6,9}
  # 차집합
  print(set_a - set_b) # {1, 2}
  # 합집합
  print(set_a | set_b) # {1, 2, 3, 6, 9}
  # 교집합
  print(set_a & set_b) # {3}
  ```

* Set 를 활용하면 다른 컨테이너에서 중복된 값을 쉽게 제거할 수 있음
  * 단, 이후 순서가 무시되므로 순서가 중요한 경우 사용할 수 없음.

### Dictionary

* key와 value가 쌍으로 이뤄진 자료구조

  * 빈 중괄호({}) 혹은 dict()을 통해 생성할 수 있음
  * key를 통해 value에 접근

  ```python
  dict_a = {'a': 1, 'b': 2}
  dict_b = {}
  dict_b['a'] = 1
  dict_b['b'] = 2
  
  # dict_a, b는 같은 내용을 담고 있음.
  ```

* key는 변경 불가능한 데이터만 활용 가능
  * string, integer, float, boolean, tuple, range

* value는 모든 값으로 설정 가능. 시퀀스가 들어가도 됨.

# 컨테이너의 특징

## 컨테이너 형 변환

|            | string |     list     |    tuple     | range |     set      | dictionary |
| :--------: | :----: | :----------: | :----------: | :---: | :----------: | :--------: |
|   string   |        |      O       |      O       |   X   |      O       |     X      |
|    list    |   O    |              |      O       |   X   |      O       |     X      |
|   tuple    |   O    |      O       |              |   X   |      O       |     X      |
|   range    |   O    |      O       |      O       |       |      O       |     X      |
|    set     |   O    |      O       |      O       |   X   |              |     X      |
| dictionary |   O    | O<br />key만 | O<br />key만 |   X   | O<br />key만 |            |

* 각자 반드시 실행해보고, 왜 그럴 수 밖에 없는지 고민해 보기

## 컨테이너 분류

* 변경 불가능한 데이터 (Immutable)
  * 리터럴 - 숫자, 문자열, 참 / 거짓(Bool)
  * range
  * tuple

* 변경 불가능한 데이터의 복사

  ```python
  a = range(10)
  b = a
  b = range(11)
  print(a) # range(10)
  ```

  * 원본이 따라 바뀌지 않음

* 변경 가능한 데이터의 복사

  ```python
  a = [1,2,3,4]
  b = a
  b[0] = 100
  print(a)  # [100, 2, 3, 4] 
  ```

  * 원본이 따라 바뀜

### 정리

* 컨테이너
  * 1. 시퀀스형
    2. 비시퀀스형
  * 1. Mutable
       * List, dictionary, set
    2. Immutable
       * string, tuple, range, ...
