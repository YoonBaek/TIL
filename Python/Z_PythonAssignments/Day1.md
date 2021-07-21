# Python 01. 데이터 & 제어문
* 학습해야할 내용
    * 기초 문법
    * 변수 및 데이터 타입

## 1. Python 예약어
* Python에서 사용할 수 없는 식별자(예약어)를 찾아 작성하시오.


```python
import keyword
print(keyword. kwlist)
```

    ['False', 'None', 'True', 'and', 'as', 'assert', 'async', 'await', 'break', 'class', 'continue', 'def', 'del', 'elif', 'else', 'except', 'finally', 'for', 'from', 'global', 'if', 'import', 'in', 'is', 'lambda', 'nonlocal', 'not', 'or', 'pass', 'raise', 'return', 'try', 'while', 'with', 'yield']


## 2. 실수 비교
* python은 부동소수점 방식을 이용하여 실수(float)를 표현하는 과정에서, 나타내고자 하는 값과의 오차가 발생하여 원하는 대로 연산 또는 비교가 되지 않을 때가 있다. 이를 참고하여, 아래와 같은 두 실수 값을 올바르게 바교하기 위한 코드를 작성하시오.


```python
num1 = 0.1 * 3
num2 = 0.3
print(num1, num2)
```

    0.30000000000000004 0.3

```python
# method1 임의의 작은 수
abs(num1 - num2) <= 1e-10
```


    True


```python
# method2 시스템 상의 machine epsilon 활용
import sys
print(abs(num1 - num2) <= sys.float_info.epsilon)
```

    True



```python
# method3 Python 3.5 이상에서 math 패키지 활용
import math
math.isclose(num1, num2)
```


    True



## 3. 이스케이프 시퀀스
* (1) 줄바꿈, (2) 탭, (3) 백슬래시를 의미하는 이스케이프 시퀀스를 작성하시오.


```python
# 1. 줄바꿈.
print("줄을 바꿉니다.\n줄을 바꿨습니다.")
```

    줄을 바꿉니다.
    줄을 바꿨습니다.

```python
# 2. 탭
print("탭으로 간격을 조절합니다.\t간격을 조절했습니다.")
```

    탭으로 간격을 조절합니다.	간격을 조절했습니다.

```python
# 3. 백슬래시
print("이것은 백슬래시 입니다. \\")
```

    이것은 백슬래시 입니다. \


## 4. String Interpolation
* '안녕, 철수야'를 string interpolation을 사용하여 출력하시오.


```python
name = "철수"
# old way (%s)
print("\'안녕, %s야\'" % name)
# recent way (format)
print("\'안녕, {}야\'".format(name))
# newest way (f-string)
print(f"\'안녕, {name}야\'")
```

    '안녕, 철수야'
    '안녕, 철수야'
    '안녕, 철수야'


## 5. 형변환
* 다음 중, 실행시 오류가 발생하는 코드를 고르시오.


```python
str(1)     # (1)
int("30")  # (2)
int(5)     # (3)
bool("50") # (4)
int("3.5") # (5)

# my_quess = (5)
```


    --------------------------------------
    
    ValueErrorTraceback (most recent call last)
    
    <ipython-input-17-d8eaaacf800f> in <module>
          3 int(5)     # (3)
          4 bool("50") # (4)
    ----> 5 int("3.5") # (5)
          6 
          7 # my_quess = (5)


    ValueError: invalid literal for int() with base 10: '3.5'


# 6. 네모 출력
* 두개의 정수 n과 m이 주어졌을 때, 가로의 길이가 n, 세로의 길이가 m인 직사각형 형태를 별(*) 문자를 이용하여 출력하시오.


```python
n, m = 5, 9
```


```python
for row in range(m) :
    print("*" * n)
```

    *****
    *****
    *****
    *****
    *****
    *****
    *****
    *****
    *****


## 7. Escape Sequence 응용
* print() 함수를 한 번만 사용하여 다음 문장을 출력하시오.
```
"파일은 c:\Windows\Users\내문서\Python에 저장이 되었습니다."
나는 생각했다. 'cd를 써서 git bash로 들어가 봐야지.'
```


```python
text = r"""
"파일은 c:\Windows\Users\내문서\Python에 저장이 되었습니다."
나는 생각했다. 'cd를 써서 git bash로 들어가 봐야지.'
"""
print(text)
```


    "파일은 c:\Windows\Users\내문서\Python에 저장이 되었습니다."
    나는 생각했다. 'cd를 써서 git bash로 들어가 봐야지.'



## 8. 근의 공식
* 다음은 이차 방정식의 근을 찾는 수식이다. 이를 파이썬 코드로 작성하시오.
$$-b \pm \sqrt{b^2-4ac} \over 2a$$


```python
import math
def quadra_formula(a, b, c) :
    x1 = (-b - math.sqrt(b**2 - 4*a*c)) / (2*a)
    x2 = (-b + math.sqrt(b**2 - 4*a*c)) / (2*a)
    if x1 > x2 :
        x1, x2 = x2, x1
    return x1, x2

a, b, c = 1, 4, -21
x1, x2 = quadra_formula(a, b, c)
print(f"The x1, x2 of {a} * x**2 + {b} * x + {c} is {x1}, {x2}")
```

    The x1, x2 of 1 * x**2 + 4 * x + -21 is -7.0, 3.0

# Python 01. 데이터 & 제어문 Workshop
* Background
    * 데이터의 입력 및 출력
    * 제어문
* Goal
    * Python 프로그래밍 언어 기본 문법의 이해
    * 데이터 입출력 방법의 이해

## Problem
## 1. 세로로 출력하기
* 자연수 number를 입력 받아, 1부터 number까지의 수를 세로로 한 줄씩 출력하시오.


```python
# input
number = int(input())
```

    10

```python
# 1 to 10 sequence for loop
for i in range(1, number + 1) :
    print(i)
```

    1
    2
    3
    4
    5
    6
    7
    8
    9
    10


## 2. 거꾸로 세로로 출력하기
* 자연수 number를 입력 받아, number부터 0까지의 수를 세로로 한 줄씩 출력하시오.


```python
# input
number = int(input())
```

    5

```python
# while loop
while number >= 0 :
    print(number)
    number -= 1
```

    5
    4
    3
    2
    1
    0


## 3. N줄 덧셈 (SWEA #2025)
* 입력으로 자연수 number가 주어질 떄, 1부터 주어진 자연수 number 까지를 모두 더한 값을 출력하시오. 단, 주어지는 숫자는 10,000을 넘지 않는다. 예를 들어, 주어진 숫자가 10일 경우 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55이므로, 출력해야 할 값은 55이다.


```python
# input
number = int(input())
```

    10

```python
total = 0
for num in range(number) :
    num += 1
    total += num
print(total)
```

    55
