# Python 02. 데이터 & 제어문 H/W
* 학습해야 할 내용
    * 데이터 타입 및 형 변환
    * 제어문

## 1. Mutable & Immutable
* 주어진 컨테이너들을 각각 변경 가능한 것(mutable)과 변경 불가능한 것(immutable)으로 분류하시오.
~~~
String, List, Tuple, Range, Set, Dictionary
~~~

1. Mutable  
    List, Dictionary, Set
2. Immutable  
    String, Tuple, Range,

## 2. 홀수만 담기
* range와 slicing을 활용하여 1부터 50까지의 숫자 중, 홀수로만 이루어진 리스트를 만드시오.


```python
print(list(range(1,51))[0::2])
```

    [1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49]


## 3. Dictionary 만들기
* 반 학생들의 정보를 이용하여 key는 이름, value는 나이인 dictionary를 만드시오.


```python
class_mates = {}
class_mates["백승윤"] = 29
class_mates["소석진"] = 27
class_mates["이채빈"] = 26

for key in class_mates :
    print(f"이름은 {key}이고, 나이는 {class_mate[key]}입니다.")
```

    이름은 백승윤이고, 나이는 29입니다.
    이름은 소석진이고, 나이는 27입니다.
    이름은 이채빈이고, 나이는 26입니다.


## 4. 반복문으로 네모 출력
* 두개의 정수 n과 m이 주어졌을 떄, 가로의 길이가 n, 세로의 길이가 m인 직사각형 형태를 별(*)문자를 이용하여 출력하시오. 단, 반복문을 사용하여 작성학시오.


```python
n, m = 5, 9
```


```python
for row in range(m) :
    for col in range(n) :
        print("*", end = "")
    print()
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


## 5. 조건 표현식
* 주어진 코드의 조건문을 조건 표현식으로 바꾸어 작성하시오.


```python
# original code
temp = 36.5
if temp >= 37.5 :
    print("입싥 불가")
else :
    print("입실 가능")
```

    입실 가능



```python
# conditional expression
print("입실 불가") if temp >= 37.5 else print("입실 가능")
```

    입실 가능


## 6. 평균 구하기
* 주어진 list에 담긴 숫자들의 평균 값을 출력하시오.


```python
scores = [80, 89, 99, 83]
```


```python
def mean(elements) :
    return sum(elements) / len(elements)

print(mean(scores))
```

    87.75


# Python 02. 데이터 & 제어문 Workshop
* Background
    * 데이터의 입력 및 출력
    * 제어문
* Goal
    * Python 프로그래밍 언어 기본 문법의 이해
    * 데이터 입출력 방법의 이해

## 1. 간단한 N의 약수 (SWEA #1933)
* 입력으로 1개의 정수 N이 주어진다. 정수 N의 약수를 오름차순으로 출력하는 프로그램을 작성하시오.
~~~
[제약 사항]
N은 1이상 1,000이하의 정수이다. (1 <= N <= 1000)
[입력]
입력으로 정수 N이 주어진다.
[출력]
정수 N의 모든 약수를 오름차순으로 출력한다.
[입력 예시]
10
[출력 예시]
1 2 5 10
~~~


```python
N = int(input())
```

    1000



```python
def get_divisor(num) :
    div_list = []
    for i in range(num) :
        i += 1
        if i ** 2 > num :
            break
        if num % i == 0 :
            div_list.append(i)
            print(i, end = " ")
    while len(div_list) > 0 :
        pop = div_list.pop()
        div = num / pop
        if div == pop :
            continue
        print(int(div), end = " ")

get_divisor(N)
```

    1 2 4 5 8 10 20 25 40 50 100 125 200 250 500 1000 

## 2. 중간값 찾기 (SWEA #2063 변형)
* 중간값은 통계집단의 수치를 크기 순으로 배열했을 때 전체의 중앙에 위치하는 수치를 뜻한다. 리스트 numbers에 입력된 숫자에서 중간값을 출력하라.
~~~
[출력 예시]
64
~~~


```python
numbers = [
    85, 72, 38, 80, 69, 65, 68, 96, 22, 49, 67,
    51, 61, 63, 87, 66, 24, 80, 83, 71, 60, 64, 
    52, 90, 60, 49, 31, 23, 99, 94, 11, 25, 24
]
```


```python
numbers.sort()
numbers[int(len(numbers) / 2)]
```




    64



## 3. 계단 만들기
* 자연수 number를 입력 받아, 아래와 같이 높이가 number인 내려가는 계단을 출력하시오.
[입력 예시]  
4  

[출력 예시]  
1<br>
1 2  
1 2 3  
1 2 3 4  


```python
for i in range(4) :
    i += 1
    for j in range(i) :
        j += 1
        print(j, end = " ")
    print()
```

    1 
    1 2 
    1 2 3 
    1 2 3 4 

