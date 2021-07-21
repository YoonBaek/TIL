# 함수

## 함수(Function)

* 특정한 기능을 하는 코드의 조각(묶음)
* 하나의 큰 프로그램을 여러 부분으로 나누어 같은 함수를 여러 상황에서 같은 함수를 여러 상황에서 호출하고 (높은 재사용성), 일부분을 수정하기 쉽다(유지보수 용이)는 장점을 가짐
* 함수의 특징
  * 함수의 이름
  * 함수의 매개변수(parameters)
  * 함수의(body) - Docstring(선택적) 및 코드셋
  * 함수의 반환값 (return)

```python
y = abs(-1)
# 1
```

여기서 abs가 함수.

## 함수를 사용해야 하는 이유

* 예를 들어, 표준 편차를 구하고 분산을 구할 경우 매번 복잡한 수식을 코드로 만들어 써야 함.

* 그러나 함수화를 통해서 아래와 같이 간소하고 반복 사용 가능한 코드 구현 가능

  ```python
  import statistics
  values = [100, 75, 85, 90, 65, 95, 90, 60, 85]
  statistics.pstdev(values)
  ```

  pstdev.\_\_doc\_\_라는 매직 메서드를 활용해서 아래와 같이 확인할 수 있다.

  ```python
  def pstdev(data, mu=None):
      """Return the square root of the population variance.
      See ``pvariance`` for arguments and other details.
      >>> pstdev([1.5, 2.5, 2.5, 2.75, 3.25, 4.75])
      0.986893273527251
      """
      var = pvariance(data, mu)
      try:
          return var.sqrt()
      except AttributeError:
          return math.sqrt(var)
  ```

## 내장함수(Built-in Functions)

* 파이썬 인터프리터에는 항상 사용할 수 있는 많은 함수와 형(type)이 내장되어 있음.

  내장함수 예시

  |        |        | 내장함수  |         |           |
  | :----: | :----: | :-------: | :-----: | :-------: |
  | all()  | abs()  | delattr() |  dir()  | setattr() |
  | ord()  | min()  |   max()   |  sum()  | complex() |
  | hash() | dict() |  help()   | eval()  |  bytes()  |
  | bin()  | bool() |   int()   | float() |    ...    |

## 함수의 선언

* 함수의 선언은 def 키워드를 활용
* 들여쓰기를 통해 함수의 body 블록을 작성함
  * doc string은 body 앞에 선택적으로 작성 가능
    * 작성 시에는 큰 따옴표 혹은 작은 따옴표 3개로 작성

* 함수에는 parameter를 넘겨줄 수 있음
* return 값을 반환 할 수 있음

## 함수의 호출

* 함수는 함수명()으로 호출
  * 매개변수가 있는 경우 함수명(매개변수1, 2, ...)로 호출

## 함수의 선언과 호출 예시

```python
a = 1
b = 2

def sum(a, b) :
  	return a + b

def diff(a, b) :
  	return a- b 

def func1(a, b) :
  	return sum(a,b) + diff(a,b)

result = func1(a, b)
print(result)
# 2
```

* pythontutor를 활용하여 작동 모습을 확인해보자

### 함수 실습문제 - 세제곱 함수

* 입력 받은 수를 세제곱 하는 함수를 작성

```python
def cube_triple(number) :
  	return number ** 3
  
print(cube(2))
# 8
```

## 함수의 리턴(return)

* 함수는 항상 반환되는 값이 있으며, 어떠한 객체라도 상관 없음

* 오직 한 개의 객체만 리턴 됨

* 복수의 객체를 리턴하는 경우

  * 복수의 객체를 하나의 튜플로 반환

  ```python
  def foo(a, b):
    	return a+b, a-b
  ```

  ```python
  type(foo(1,2))
  # tuple
  ```

* 명시적인 리턴 값이 없는 경우

  * None을 반환 함

  ```python
  def hello() :
    	print("hello")
      
  type(hello())
  # None
  ```

### 함수 실습 문제 - 사각형의 넓이

* 너비와 높이를 입력 받아 사각형의 넓이와 둘레를 tuple로 반환하는 함수 rectangle을 작성하시오.

  ```python
  def rectangle(width, height) :
    	return width*height, (width + height) * 2
    
  print(rectangle(2,4))
  # 8, 12
  ```

# 함수 input

## 위치 인자(Positional Arguments)

* 기본적으로 함수 호출시 인자는 위치에 따라 함수 내에 전달됨.

  ```python
  def add(x, y) :
    	return x + y
  add(2,3) # x, y = 2,3
  # 5
  ```

  * 매개변수
    * 함수에 입력으로 전달된 값을 받는 변수
      * 위에서는 x, y

  * argument
    * 함수를 입력할 때 전달 받는 입력값
      * 위에서는 2, 3

## 기본 인자값(Default Arguments Values)

* 기본 값을 지정하여 함수 호출 시 인자 값을 설정하지 않도록 함

  * 정의된 것 보다 더 적은 개수의 인자들로 호출 될 수 있음
    ```python
    def add(x, y = 0) :
        return x + y
    add(2) # x, y = 2, 0
    # 2
    ```

## 키워드 인자 (Keyword Arguments)

* 직접 변수의 이름으로 특정 인자를 전달할 수 있음

* 키워드 인자 다음에 위치 인자를 활용할 수 없음

    ```python
    def add(x, y) :
      	return x + y
    add(y = 2,x = 3) # x, y = 3, 2
    # 5
    ```

## 정해지지 않은 여러개의 인자 처리

* 함수가 임의의 개수 인자로 호출될 수 있도록 지정

* 인자들은 튜플로 묶여 (packing)처리 되며, 매개변수에 *을 붙여 표현

  ```python
  def add(*args) : # 관습적으로 args라는 이름을 사용.
    	total = 0
    	for arg in args :
        	total += arg
      return total
  add(2,3,4,5)
  # 14
  ```

## 가변 키워드 인자 (Arbitrary Keyword Arguments)

* 함수가 임의의 개수 인자를 키워드 인자로 호출 할 수 있도록 지정
* 인자들은 **딕셔너리로 묶여** 처리되며, 매개변수 **를 붙여 표현

	  ```python
  def family(**kwargs) : # 별이 두 개 인건 두번 푼다는 의미.
    	for key, value in kwargs.items :
        	print(f"{key} : {value}")
  famil(father="Baek", son="Baekß")
  # father : Baek
  # son : Baek

## 함수 정의 주의 사항

* 기본 인자 값을 가지는 인자 다음에 기본 값이 없는 인자로 정의할 수 없음.

  ```python
  add(x = 3, 5)
  # SyntacError : positional argument follows keyword argument
  ```

  ```python
  add(*args, x)
  add(1, 2, 3)
  # TypeError : 가변 인자 리스트가 위치 인자보다 앞쪽에 올 수 없음
  ```

* 가장 좋은 구조

  ```python
  my_info(x, y, *args, **kwargs)
  ```

# 스코프(Scope)

## 함수 스코프(scope)

* 함수는 코드 내부에 지역스코프(local scope)를 생성하며, 그 외의 공간인 전역 스코프(global scope)로 구분
* 스코프
  * 전역 스코프(global scope) : 코드 어디에서든 참조할 수 있는 공간
  * 지역 스코프(local scope) : 함수가 만든 스코프. 함수 내부에서만 참조 가능
* 변수
  * 전역 변수(global variable) : 전역 스코프 내에서 지정된 변수
  * 지역 변수(local variable) : 지역 스코프 내에서 지정된 변수

## 변수 수명주기(life cycle)

* 변수는 각자의 수명 주기가 존재
  * 빌트인 스코프(built-in scope)
    * 파이썬이 실행된 이후부터 영원히 유지
  * 전역 스코프(global scope)
    * 모듈이 호출된 시점 이후 혹은 인터프리터가 끝날 때 까지 유지
  * 지역(함수) 스코프(local scope)
    * 함수가 호출될 때 생성되고, 종료될 때 까지 유지

## 함수 스코프 예시

```python
def func() :
  	a = 20
    print('local', a)
func()
print('global', a)

# NameError: name 'a' is not defined
# 로컬 a만 지정되고, 함수가 끝나고 사라지기 때문에 글로벌 스코프에서는 a가 존재하지 않음.
```

* python tutor를 통해서 확인해보자

## 이름 검색 규칙(Name Resolution)

* 파이썬에서 사용되는 이름(식별자)들은 이름공간(namespace)에 저장되어 있음
* 아래와 같은 순서로 이름을 찾아나가며, **LEGB Rule**이라고 부름
  * Local scope : 함수
  * Enclosed scope : 특정 함수의 상위 함수
  * Global scope : 함수 밖의 변수, Import 모듈
  * Built-in scope : 파이썬 안에 내장되어 있는 함수 또는 속성
* **즉, 함수 내에서는 바깥 스코프의 변수에 접근 가능하나 수정은 할 수 없음.**

## LEGB 예시

```python
a = 0
b = 1
def enclosed() :
		a = 10
    c = 3
    def local(c) :
      	print(a, b, c)
    local(300)
    print(a, b, c)
enclosed()
print(a, b)

# 10, 1, 300
# 10, 1 ,3
# 0, 1
```

## global

* 현재 코드블럭 전체에 적용되며, 나열된 식별자(이름)들이 전역 변수 임을 나타냄

  * global에 나타난 이름들은 같은 코드 블럭 내에서 global 앞에 등장할 수 없음

  * global에 나열된 이름은 매개변수, for 루프 대상, 클래스/함수 정의 등으로 정의되지 않아야 함

    ```python
    a = 10
    def func1() :
      	global a
        a = 3
    print(a)
    func1()
    print(a)
    # 10 3
    ```

## global 관련 에러

```python
a = 10
def func1() :
  	print(a)
  	global a
    a = 3
print(a)
func1()
print(a)
# SyntaxError: name 'a' is used prior to global declaration
```

```python
a = 10
def func1(a) :
  	global a
    a = 3
print(a)
func1(3)
print(a)
# SyntaxError: name 'a' is parameter and global
```

## nonlocal

* 전역을 제외하고 가장 가까운 (둘러싸고 있는) 스코프의 변수를 연결하고자 함
  *  non local에 나열된 이름은 같은 코드블록에서 nonlocal 앞에 등장할 수 없음
  * nonlocal에 나열된 이름은 매개변수, for 루프 대상, 클래스/함수 정의 등으로 정의 되지 않아야 함

* global과는 달리 이미 존재하고 있는 이름과만 연결 가능.

## non local 예시

```python
# non local 예시
x = 0
def func1() :
  	x = 1
  	def func2() :
    		nonlocal x
      	x = 2
    func(2)
    print(x)
func1()
print(x)
# 2 
# 0
```

## nonlocal, global 비교

* nonlocal 이름 공간 상에 존재하는 변수만 가능

## 주의사항

* 기본적으로 함수에서 선언된 변수는 Local scope에서 생성되며, 함수 종료시 사라짐
* 해당 스코프에 변수가 없는 경우 LEGB rule에 의해 이름을 검색함
  * 변수에 접근은 가능하지만, 해당 변수를 수정할 수는 없음
  * 값을 할당하는 경우 해당 스코프의 이름공간에 새롭게 생성되기 때문
  * **단, 함수 내에서 필요한 상위 스코프 변수는 인자로 넘겨서 활용할 것(클로저* 제외)**
    * (참고) 클로저란? 어떤 함수 내부에 중첩된 형태로써 외부 스코프 변수에 접근 가능한 함수
* 상위 스코프에 있는 변수를 수정하고 싶다면 global, nonlocal 키워드를 활용 가능
  * 단, 코드가 복잡해지면서 변수의 변경을 추적하기 어렵고, 예기치 못한 오류가 발생
  * 가급적 사용하지 않는 것을 권장하며, **함수로 값을 바꾸고자 한다면 항상 인자로 넘기고 리턴 값을 사용하는 것을 추천**

# 재귀함수

## 재귀 함수(recursive function)

* 자기 자신을 호출하는 함수
* 무한한 호출을 목표로 하는 것이 아니며, 알고리즘 설계 및 구현에서 유용하게 활용
  * 알고리즘 중 재귀 함수로 로직을 표현하기 쉬운 경우가 있음(예 - 점화식)
  * 변수의 사용이 줄어들며, 코드의 가독성이 높아짐
* 1개 이상의 base case(종료되는 상황)가 존재하고, 수렴하도록 작성
  * 같은 문제를 다른 Input 값을 통해서 해결하는 과정
    * 큰 문제를 해결하기 위해 작은 문제로 좁히고, 작은 문제의 해답을 이용하여 해결
  * 작은 문제는 base case에 도달하여 재귀함수가 끝날 수 있도록 함

### 재귀 함수 실습 문제

* 팩토리얼 구현

  ```python
  def factorial(n) :
      if n > 1 :
          return n * factorial(n-1)
      if n == 1 :
          return n
  print(factorial(6))
  ```

  * 반복문
    * n이 1보다 큰 경우 반복문 실행, n은 1씩 감소
    * 마지막에 n이 1이면 더 이상 반복문을 돌지 않는다.
  * 재귀함수
    * 재귀 함수를 호출하며, n은 1씩 감소
    * 마지막에 n이 1이면 더 이상 재귀함수를 돌지 않음.

## 재귀 함수 주의사항

* 재귀함수는 base case에 도달할 때 까지 함수를 호출함
* 메모리스택이 넘치게 되면(stack overflow)프로그램이 동작하지 않게 됨
* 파이썬에서는 최대 재귀깊이가 1,000번으로, 이를 넘어가게 되면 RecursionError발생

### 재귀함수 실습문제2

* 피보나치 수열을 반복문과 재귀함수로 구현

  ```python
  # 재귀
  def fibo(n) :
      if n <= 2 :
          return 1
      else :
          return fibo(n-2) + fibo(n-1)
        
  # 함수 호출이 기하급수적으로 많아지므로 좋은 방법은 아니다.
  ```

  ```python
  # 반복문
  # 재미있게도 파이썬 메인 홈에 나와있는 수식이다.
  def fib(n):
      a, b = 0, 1
      while a < n:
          print(a, end=' ')
          # 파이썬의 swap 기능을 활용
          a, b = b, a+b
      print()
  fib(1000)
  # 기하급수 적인 호출 없이 빠르게 쓸 수 있다.
  ```
