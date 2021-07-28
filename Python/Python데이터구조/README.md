# 데이터 구조

## 데이터 구조(Data Structurers)

* Niklaus Wirth
  * Algoriths + DataStructures = Programs
  * 프로그래밍 언어 파스칼 설계(1969)
  * 1984 튜링상 수상
* 데이터에 편리하게 접근하고, 변경하기 위해 데이터를 저장하거나 조작하는 법
  * '자료구조' 라고도 함
* 순서가 있는 데이터 구조
* 순서가 없는 데이터 구조 등으로 구성

## 문자열

* 문자들의 나열(sequence of characters)

* 문자열의 특징

  * 변경할 수 없고(immutable)

    ```python
    a = 'my string'
    a[-1] = 'f' # TypeError
    ```

  * 순서가 있고(ordered)

  * 순회 가능한(iterable)

### 문자열 인덱스와 슬라이싱

* 파이썬에서 문자열은 아래와 같이 접근할 수 있음

  ```python
  a = "abcdefghi"
  a[9] # IndexError
  ```

  |       |  a   |  b   |  c   |  d   |  e   |  f   |  g   |  h   |  i   |
  | :---: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
  | index |  0   |  1   |  2   |  3   |  4   |  5   |  6   |  7   |  8   |
  | index |  -9  |  -8  |  -7  |  -6  |  -5  |  -4  |  -3  |  -2  |  -1  |

  * s[2:5] => "cde"
  * s[-6:-2] => "defg"

  * s[2:5:2] => "ce"
  * s[-6:-1:3] =>  "dg"
  * s[2:5:-1] => ""
  * s[5:2:-1] => "fed"

* 독특한 기능
  * s[::] => "abcdefghi"
    * s[1:len(s):1] 과 동일
  * s[::-1] : 뒤집기에 활용
    * s[-1:-len(s):-1] 과 동일

### 문자열 메서드

* 메서드를 전부 외울 필요는 없으나 이 정도는 알아두면 좋음

#### .find(x)

* "x"의 첫번째 위치를 반환. 없으면 -1을 반환

  ```python
  "apple".find("p") # 1
  "apple".find("k") # -1
  ```

#### .index(x)

* find와 같이 동작하나 없는 문자열 찾을 때 ValueError

  ```python
  "apple".index("p") # 1
  "apple".index("k") # ValueError
  ```

#### .replace(old, new[, count])

* 대괄호 -> 선택적 인자를 의미 (doc)

* 바꿀 대상 글자를 새로운 글자로 바꿔서 반환 (복사본 반환)

* count를 지정하면, 앞의 count 개만큼 바꿔서 시행

  ```python
  "woooooWoo".replace("o", "!", 2) # w!!oooWoo
  ```

#### .strip([chars])

* 특정한 문자들을 지정하면

  * 왼쪽을 제거하거나(lstrip), 오른쪽을 제거하거나(rstrip), 양쪽을 제거(strip)

* 문자열을 지정하지 않을 수도 있는데, 지정하지 않으면 공백을 제거함

  ```python
  " dirtydata\n ".strip() # "dirtydata"
  ```

#### .split(sep=None)

* 문자열을 특정한 단위로 나눠서 리스트로 반환

  ```python
  "a,b,c".split("_")
  # ["a,b,c"]
  "a b c".split()
  # ["a", "b", "c"]
  ```

#### 'seperator'.join(iterable)

* 반복 가능한(iterable) 컨테이너 요소들을 seperator(구분자)로 합쳐 문자열 반환

  ```python
  "-".join(["010","1234","1234"])
  # 010-1234-1234
  ```

#### 나머지 메서드들

* **capitalize()** : 첫문자를 대문자, 나머지는 소문자
* **title()** : ''나 공백 이후의 단어 첫 문자를 대문자로
* **upper()** : 모두 대문자
* **lower()** : 모두 소문자로
* **swapcase()** : 소문자는 대문자로 대문자는 소문자로

#### 검증 메서드

* **.isalpha()** : 알파벳 문자 여부
  * 단순 알파벳이 아닌 유니코드상 letter (한국어도 포함)
* **.isupper()** : 대문자 여부
* **.islower()** : 소문자 여부
* 기타 다양한 검증 메서드 들이 있지만 암기보다는 필요시 확인하고 사용

## 리스트

### 리스트

* 순서가 있는 시퀀스, 인덱스로 접근
* 리스트의 특징
  * 변경 가능하고(Mutable)
  * 순서가 있고(Ordered)
  * 순회 가능한(Iterable)

### 리스트 메서드

#### .append(x)

* 리스트 끝에 값을 추가함

  ```python
  cafe = ['starbucks', 'twosomeplace', 'hollys']
  cafe.append('coffeebean')
  ```

#### .extend(iterable)

* 리스트에 iterable의 각 항목을 추가함

    ```python
   cafe.extend("gongcha")
    # ['starbucks', 'twosomeplace', 'hollys', 'coffeebean', 'g', 'o', 'n', 'g', 'c', 'h', 'a']
    ```

### .insert(i, x)

* 정해진 위치에 값을 추가

### .remove(x)

* 리스트에서 값이 x인 첫번째 항목 삭제

  ```python
  numbers = [1, 2, 3, "hi"]
  numbers.remove("hi") # [1, 2, 3]

### .pop(i)

* 정해진 위치 i에 있는 값을 삭제하고, 그 항목을 반환함

* i가 지정되지 않으면, 마지막 항목을 삭제하고 반환함

  * stack의 개념이 적용됨

  ```python
  numbers = [1,2,3,"hi"]
  print(numbers.pop(), numbers)
  
  # "hi" [ 1, 2, 3]
  ```

### .clear()

* 모든 항목을 삭제함

#### .index(x)

* 첫번째 x 값을 찾아 해당 index 값을 반환
* string의 index와 동일
  * 없으면 ValueError를 일으키는 것도 동일

#### .count(x)

* 원하는 값 x의 개수를 반환함

  ```python
  numbers = [1, 2, 3, 1, 1]
  numbers.count(1) # 3
  ```

#### .sort()

* 원본 리스트를 정렬함. None 반환 (복사가 아닌 원본 변경)

* sorted 함수와 비교할 것

  ```python
  numbers = [3, 2, 5 ,1]
  result = numbers.sort()
  print(numbers, result())
  # [1, 2, 3, 5] None
  ```

  ```python
  numbers = [3, 2, 4, 1]
  result = sorted(numbers)
  print(numbers, result())
  # [3, 2, 4, 1] [1, 2, 3, 4]
  ```

#### .reverse()

* 원본 리스트를 뒤집음
* 순서가 바뀌는 것이 아님.

### 리스트 복사

* 리스트 복사 확인하기

  ```python
  original_list = [1, 2, 3]
  copy_list = original_list
  copy_list[0] = "hello"
  print(original_list, copy_list)
  # ["hello", 2, 3] ["hello", 2, 3]
  ```

### 리스트 복사 - 얕은 복사(shallow copy) 1

* Slice 연산자를 활용하여 같은 원소를 가진 리스트지만 연산된 결과를 다른 주소에 복사

  ```python
  copy_list = original_list[:]
  copy_list[0] = "hello"
  print(original_list, copy_list)
  # [1, 2, 3] ["hello", 2, 3]
  ```

### 리스트 복사 - 얕은 복사(shallow copy) 2

* list() 활용하여 같은 원소를 가진 리스트지만 연산된 결과를 복사 (다른 주소)

  ```python
  copy_list = list(original_list)
  copy_list[0] = "hello"
  print(original_list, copy_list)
  # [1, 2, 3] ["hello", 2, 3]
  ```

### 리스트 복사 주의사항

* 복사하는 원소가 리스트인 경우까지는 복사가 안됨

* deepcopy 활용

  ```python
  import copy
  a = [1, 2, [1, 2]]
  b = copy.deepcopy(a)
  print(a, b)
  b[2][0] = 0
  print(a, b)
  # [1, 2, [1, 2]] [1, 2, [1, 2]]
  # [1, 2, [1, 2]] [1, 2, [0, 2]]
  ```

### List comprehension

* 표현식과 제어문을 통해 특정한 값을 가진 리스트를 생성하는 방법

  [<expression>for <변수> in <iterable> if <조건식>]

  ```python
  # 1~3의 세제곱 리스트 만들기
  cubic_list = []
  for number in range(1, 4) :
      cubic_list.append(numbers ** 3)
  cubic_list
  # [1, 8, 27]
  ```

  ```python
  # 위 코드와 정확히 같은 코드
  [number**3 for number in range(1,4)]
  ```

  ```python
  # 짝수만 담긴 리스트 생성
  [num for num in range(1, 4) if num % 2 == 0]
  ```

#### 실습

* 다음의 반복문을 list comprehension으로 표현하시오

  ```python
  girls = ['jane', 'ashley']
  boys = ['justin', 'eric']
  ```

  ```python
  pair = []
  for boy in boys :
      for girl in girls :
          pair.append((boy, girl))
  ```

* lc

  ```python
  [(boy, girl) for boy in boys for girl in girls] 
  ```

* 코드는 간단해지나 가독성 측면에서 과연 lc가 좋을지는 생각을 해보고 코드를 짜야한다.

* 두가지 스타일 모두로 코드를 작성할 줄 알아야

### Built-in Function - map

* map(function, iterable)

* 순회 가능한 데이터 구조(iterable)의 모든 요소에 함수(function)적용하고, 그 결과를 map object로 반환

  ```python
  numbers = [1, 2, 3]
  result = map(str, numbers)
  print(result, type(result))
  # <map object at 0x10e2ca100> <class map>
  list(result)
  # ['1', '2', '3']
  ```

#### 활용 사례

* 알고리즘 문제 풀이시 input 값들을 숫자로 바로 활용하고 싶을 때

  ```python
  n, m = map(int, input().split())
  
  print(n, m)
  print(type(n), type(m))
  # 3, 5
  # int, int
  ```

### Built-in Function -filter

* filter(function, iterable)

* 순회 가능한 데이터 구조(iterable)의 모든 요소에 함수(function)적용하고, 그 결과가 True인 것들을 filter object로 반환

  ```python
  def ood(n) :
      return n % 2
  numbers = [1, 2, 3]
  result = filter(odd, numbers)
  print(list(result), type(result))
  # [1, 3] <class 'filter'>
  ```

### Built-in Function - zip

* 복수의 iterable을 모아 튜플을 원소로 하는 zip object를 반환

  ```python
  girls = ['jane', 'ashley']
  boys = ['justin', 'eric']
  pair = zip(girls, boys)
  print(list(pair), type(pair))
  # [('jane', 'ashley'), ('justin', 'eric')], class 'zip'
  ```

## Set

#### .add(elem)

* 세트에 값을 추가
* 추가한다고 해서 순서가 생기는 것은 아님

#### .update(*others)

* 여러 값을 추가
* 중복이 있을 시 중복을 제외하고 추가

#### .remove(elem)

* 세트에서 삭제하고, 없으면 KeyError

#### .discard(elem)

* 삭제하는 것이 없어도 에러가 발생하지 않음

#### .pop()

* 임의의 원소를 제거해 반환. 세트가 비어있는 경우 KeyError

  ```python
  a = {"사과", "바나나", "수박"}
  print(a.pop)
  # 바나나
  print(a)
  # {사과, 바나나}
  ```

## Dictionary

* 키와 밸류로 구성된 데이터 구조
* 딕셔너리의 특징
  * Mutable
  * Unordered
  * Iterable

### 메서드들

#### .get(key [,default])

* key에 대응하는 value를 가져옴

* key에 딕셔너리가 없어도 keyError가 발생하지 않으며, default를 돌려줌 (기본 : None)

  ```python
  mydict = {"apple" : "사과"}
  print(my_dict.get("pineapple"))    # None
  print(my_dict.get("pineapple", 0)) # 0
  ```

#### .pop(key[,default])

* key가 딕셔너리에 있으면 제거하고 해당값을 반환
* 역시 default 값을 통해 해당 키가 없을 때 출력을 정할 수 있음.

#### .update()

* 값을 제공하는 key, value로 갱신(기존 key는 덮어씀.)

  ```python
  ```

  

