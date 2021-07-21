# Python 03. 함수 (hw)
* 학습해야 할 내용
    * 함수의 선언과 호출
    * 함수의 매개변수와 인자

## 1. Built-in 함수
* 파이썬에서 기본으로 사용할 수 있는 built-in 함수를 최소 5가지 이상 작성하시오.
    * sum(), pow(), dir(), help(), round(), int(), float(), str(), type(), bool(), list(), set(), tuple(), "string".replace() ...

## 2. 정중앙 문자

* 문자열을 전달 받아 해당 문자열의 정중앙 문자를 반환하는 get_middle_char 함수를 작성하시오. 단, 문자열의 길이가 짝수일 경우에는 정중앙 문자 2개를 반환한다.


```python
def get_middle_char(word) :
    target_idx = int(len(word) / 2)
    target = word[target_idx]
    if len(word) % 2 == 0 :
        target = word[target_idx-1] + target
    return target

print(get_middle_char("ssafy"))
print(get_middle_char("coding"))
```

    a
    di


## 3. 위치 인자와 키워드 인자
* 다음과 같이 함수가 선언되어 있을 때, 보기 (1)~(4) 중에서 실행시 오류가 발생하는 코드를 고르시오.


```python
def ssafy(name, location="서울"):
    print(f"{name}의 지역은 {location}입니다.")

# (1)
ssafy("허준") # location을 디폴트 값으로 파싱해서 정상 출력
# (2)
ssafy(location="대전", name="철수") # 파라미터를 전달할 경우 순서와 무관하게 출력
# (3)
ssafy("영희", location = "광주") # 포지션에 의해 "영희"는 name으로 해석 된다.
# (4)
ssafy(name="길동", "구미") # 파라미터를 지정해줄경우 그 순간 이후로는 위치로 파라미터를 보지 않으므로 에러가 남.
```


      File "<ipython-input-13-03af08a88b78>", line 11
        ssafy(name="길동", "구미")
                             ^
    SyntaxError: positional argument follows keyword argument



## 4. 나의 반환값은
* 다음과 같이 함수를 선언하고 호출하였을 때, 변수 result에 저장된 값을 작성하시오.


```python
def my_func(a, b) :
    c = a + b
    print(c)

# return 값이 없는 함수는 None을 반환하기 때문에 result에 저장된 값은 None이다.
result = my_func(3,7)
print(result)
```

    10
    None


## 5. 가변 인자 리스트
* 가변 인자 리스트를 사용하여, 갯수가 정해지지 않은 여러 정수들을 전달 받아 해당 정수들의 평균 값을 반환하는 my_avg 함수를 작성하시오.


```python
def my_avg(*args) :
    total = 0
    length = 0
    for arg in args :
        length += 1
        total += arg
        
    return round(total/length, 1)

print(my_avg(77, 83, 95, 80, 70))
```

    81.0

# Python 03. 함수 (workshop)

* Background
  * 함수
  * 매개변수와 인자
* Goal
  * 함수의 선언과 호출에 대한 이해
  * 함수의 반환에 대한 이해

## Problem

## 1. List의 합 구하기

* 정수로만 이루어진 list를 전달받아 해당 list의 모든 요소들의 합을 반환하는 list_sum 함수를 built-in 함수인 sum() 함수를 사용하지 않고 작성하시오.


```python
def list_sum(args) :
    total = 0
    for arg in args :
        total += arg
    return total

list_sum([1,2,3,4,5])
```


    15



## 2. Dictionary로 이뤄진 List의 합 구하기

* Dictionary로 이루어진 list를 전달 받아 모든 dictionary의 'age' key에 해당하는 value들의 합을 반환하는 dict_list_sum 함수를 built-in 함수인 sum() 함수를 사용하지 않고 작성하시오.


```python
def dict_list_sum(dicts):
    total = 0
    for d in dicts :
        total += d['age']
    return total

dict_list_sum([{'name':'kim', 'age':12},
              {'name':'lee', 'age': 4}])
```


    16



## 3. 2차원 List의 전체 합 구하기

* 정수로만 이루어진 2차원 list를 전달받아 해당 list의 모든 요소들의 합을 반환하는 all_list_sum 함수를 built-in 함수인 sum() 함수를 사용하지 않고 작성하시오.


```python
def all_list_sum(arr) :
    total = 0
    for l in arr :
        for i in l :
            total += i
    return total

arr = [
    [1], 
    [2, 3], 
    [4, 5, 6], 
    [7, 8, 9, 10],
]

all_list_sum(arr)
```


    55



## Reduce

* 반드시 인자로 함수(함수() 아님)를 넣어줘야 함.
* 뒤쪽 컨테이너에서 원소를 차례로 꺼내서 함수를 연산


```python
from functools import reduce

reduce(lambda x, y : x + y, reduce(lambda x, y : x + y, arr))
```


    55
