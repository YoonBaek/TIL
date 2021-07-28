# OOP

* 객체
* OOP
* 클래스와 인스턴스
* 상속

# 객체

* 세상에 존재하는 설명 가능한 모든 것. 행동요소, 상태, 속성 등
* 현실 세계에 존재하는 모든 것을 객체로서 표현 -> 객체지향
* 파이썬의 모든 것. 변수, 함수, 클래스 ...

## 객체

* 클래스와 인스턴스로 이루어져 있음
* 실세계의 비유로 설명하자면, 
  * 집의 청사진이나 축소 모형 따위는 실 거주하는데 필요없다
  * 실제 필요한 것은 집
  * 여기서 청사진은 클래스를, 실제 집은 객체를 나타낸다.

* my_lower("Hi") vs "Hi".lower()
  * 전자
    * 함수가 객체를 처리
  * 후자
    * 객체가 메서드를 호출

> 객체는 특정 타입의 인스턴스다

* 123, 900, 5는 모두 int의 인스턴스
* 'hello', 'bye'는 모두 string의 인스턴스
* [232. 89,1].[]은 모두 list의 인스턴스

## 객체의 특징

* 타입(type) : 어떤 연산자와 조작 가능한가?

  ```python
  print(type('1'))
  # <class 'str'>
  ```

  string도 클래스다. string의 특징을 상속하면 string 연산자와도 조작 가능할 듯

* 속성(attribute) : 어떤 상태(데이터)를 가지는가?

* 조작법(method) : 어떤 행위(함수)를 할 수 있는가?

## is연산자

* is

  * 객체의 아이덴티티를 검사하는 연산자

  ```python
  type('a') is int # False
  ```

## isinstance 함수

* isinstance(object, classinfo)

  * classinfo의 instance거나 subclass*인 경우 Trye
  * classinfo가 tuple인 경우 (type으로 구성된) 하나라도 일치하면 True
  * classinfo가 type이거나 type으로 구성되지 않은 경우 TypeError

  ``` python
  isinstance(10, (bool, int)) # True
  isinstance(0, (bool, "hi", complex)) # False가 아닌 TypeError 
  ```

## 객체 - 속성(attribute)

* \<object\>. \<attribute\>

* 속성은 객체의 상태/데이터

* 호출하는 개념은 아님.

  ```python
  (3+4j).real # 3.0
  ```

## 객체 - 메서드

* \<object\>. \<method\>

* 메서드는 특정 객체에 적용될 수 있는 행위를 뜻하며, 일반적으로 클래스에 정의된 함수

  ```python
  "yoon".capitalize() # Yoon
  {'a': 'apple'}.items() # dict_items([('a', 'apple')])
  ```

# 객체지향 프로그래밍

## 프로그래밍 패러다임

* 기능에 따라 프로그래밍 언어를 분류하는 방법
  1. 명령형 프로그래밍(프로그래머가 기계에게 상태를 변경하는 방법을 지시
     1. 절차지향형 프로그래밍
     2. 객체지향형 프로그래밍

## 객체 지향 프로그래밍 

* 객체 지향 프로그래밍(Object Oriented Programming)은 컴퓨터 프로그래밍의 패러다임의 하나이다.
* 객체 지향 프로그래밍은 컴퓨터 프로그램을 명령어의 목록으로 보는 시각에서 벗어나 여러개의 독립된 단위, 즉 "객체"들의 모임으로 파악하고자 하는 것이다.

### 절차 지향(절차적) 프로그래망(Procedural Programming)

* 데이터 보다는 함수 지향적

### 객체 지향 프로그래밍

* 데이터와 기능 분리, 추상화된 구조(인터페이스)

## Why 객체 지향 프로그래밍?

* 현실 세계를 프로그램 설계에 반영(추상화)

```python
class Person:
    def __init__(self, name, gender):
        self.name = name
        self.gender = gender
        
    def greeting(self):
        print(f"안녕하세오, {self.name}입니다.")
```

```python
jimin = Person("지민", "남")
jimin.greeting()
```

```
안녕하세요, 지민입니다.
```

```python
jieun = Person("아이유", "야")
jieun.greeting()
```

```
안녕하세요, 아이유입니다.
```

추상화해서 인스턴스를 찍어낼 수 있음

## 객체 지향형 프로그래밍의 일반적인 모습

* 사각형 넓이 구하기 코드

### 절차 지향 프로그래밍

```python
# 함수를 호출하는 방식으로 진행
def area(x, y):
    return x * y
    
def circumference(x,y) :
    return 2 * (x + y)

a, b, c, d = 10, 30, 300, 20

square_area1 = area(a, b)
square_circumference1 = circumference(a, b)
square_area2 = area(c, d)
square_circumference2 = circumference(c, d)
```

### 객체 지향 프로그래밍

```python
class Rectangle : # 공통점을 통해 추상화
    def __init__(self, x, y):
        self.x = x
        self.y = y
    def area(self):
        return self.x * self.y
    
    def circumference(self):
        return 2 * (self.x + self.y)
    
r1 = Rectangle(10, 30)
r1.area()
r1.circumference()

r2 = Rectangle(300, 20)
r2.area()
r1.circumference()

# r1, r2 ==> 인스턴스
```

## 객체 지향 프로그래밍
* 사각형 - 클래스(class)
* 사각형의 정보 - 속성(attribute)
    * 가로 길이, 세로 길이
* 사각형의 행동 - 메서드(method)
    * 넓이, 둘레

## 클래스(class)와 인스턴스(instance)

* 클래스 정의


```python
class MyClass:
    pass
```

* 인스턴스 생성


```python
my_instance = MyClass()
```

* 메서드


```python
my_instance.my_method()
```

* 속성


```python
my_instance.my_attribute
```

* 클래스를 정의하고, 인스턴스 들을 만들어 활용함
    * 클래스: 객체들의 분류(class)
    * 인스턴스: 하나하나의 실체/예(instance)


```python
# class의 타입
# 메타클래스라 불리는 type이라는 클래스가 따로 있음
class Person:
    pass
type(Person)
```




    type




```python
person1 = Person()
```


```python
isinstance(person1, Person)
```




    True




```python
type(person1)
```




    __main__.Person



## 속성
* 특정 데이터 타입 / 클래스 객체들이 가지게 될 상태 / 데이터를 의미


```python
class Person:
    def __init__(self, name):
        self.name = name
```


```python
person1 = Person("지민")
```


```python
person1.name
```




    '지민'



## 메서드
* 특정 데이터 타입 / 클래스의 객체에 공통적으로 적용 가능한 행위(함수)


```python
class Person:
    
    def talk(self):
        print("안녕")
        
    def eat(self, food):
        print(f"{food}를 냠냠")
```


```python
person1 = Person()
```


```python
person1.talk()
```

    안녕

```python
person1.eat("피자")
```

    피자를 냠냠


## self
* 인스턴스 자기 자신
* 파이썬에서 인스턴스 메서드는 호출 시 첫번째 인자로 인스턴스 자신이 전달되게 설계 (반드시!)
    * 매개변수 첫 번쨰 인자로 self를 정의

## 생성자(constructor)
* 인스턴스 객체가 생성될 때 호출되는 메서드


```python
class Person:
    def __init__(self, name):
        print(f"instance {name} init")
```


```python
person1 = Person("yoonbaek")
```

    instance yoonbaek init


## 소멸자(destructor)
* 인스턴스 객체가 소명되기 직전에 호출되는 메서드


```python
class Person:
    def __del__(self):
        print(f"instance del")
```


```python
person1 = Person()
del person1
```

    instance del


## 매직 메서드
* double underscore(\_\_)가 있는 메서드는 특수한 동작을 위해 만들어진 메서드로, 스페셜 메서드 혹은 매직 메서드라고 불림
* 예시
    * \_\_str(self)\_\_, \_\_len(self)\_\_
    * \_\_gt(self, other)\_\_, \_\_le\_\_(self, other) ...
* 객체의 특수 조작 행위를 지정(함수, 연산자 등
    * \_\_str\_\_ : 해당 객체의 출력 형태를 지정
    * \_\_gt\_\_ : 부등호 연산자
* dir을 통해서 확인 가능

# 클래스와 인스턴스

## 인스턴스 변수
* 인스턴스의 속성(aㅡttribute)
* 메서드에서 self.<name>으로 정의
* 인스턴스가 생성된 이후 <instancename>.<name> 등오로 활용


```python
class Person:
    def __init__(self, name):
        self.name = name
```


```python
john = Person("john")
print(john.name)
```

    john

```python
# 새롭게 할당도 가능
john.name = "John Kim"
```


```python
print(john.name)
```

    John Kim


## 클래스 변수
* 클래스 속성(attribute)
* 클래스 속성(attribute)
* 모든 인스턴스가 공유
* 클래스 선언 내부에서 정의
* <classname>.<name>으로 접근 및 할당
    * 다만, 인스턴스도 접근 가능


```python
class Circle:
    pi = 3.14
```


```python
c1 = Circle()
c2 = Circle()
```


```python
# 클래스 변수는 모든 변수가 공유
print(Circle.pi)
print(c1.pi)
print(c2.pi)
```

    3.14
    3.14
    3.14


인스턴스 자기 자신에서 찾음 -> 없음 -> 클래스에서 찾음

## 인스턴스와 클래스 간의 이름 공간(namespace)
* 클래스를 정의하면, 클래스에 해당하는 이름 공간 생성
* 인스턴스를 만들면, 인스턴스 객체가 생성되고 이름 공간 생성
* 인스턴스에서 특정 속성에 접근하면, 인스턴스-클래스 순으로 탐색


```python
# Person 정의
class Person:
    name = "unknown"
    
    def talk(self):
        print(self.name)
```


```python
p1 = Person()
p1.talk()
```

    unknown

```python
# p2 인스턴스 변수 설정 후
p2 = Person()
p2.talk()
p2.name = "Kim"
p2.talk()
```

    unknown
    Kim


* 인스턴스 변수를 만드는 순간 인스턴스 변수가 먼저 나옴
* 다만 이건 덮어쓴 구조인데, 좋은 구조는 아님.

## 메서드의 종류
* 인스턴스 메서드
* 클래스 메서드
* 스태틱 메서드

## 인스턴스 메서드
* 인스턴스가 사용할 메서드
* 클래스 내부에 정의되는 메서드의 기본
* 호출 시, 첫번쨰 인자로 인스턴스 자기자신(self)이 전달됨


```python
class MyClass:
    def instance_method(self) :
        pass
```

## 클래스 메서드
* 클래스가 사용할 메서드
* @classmethod 데코레이터를 사용하여 정의
* 호출 시, 첫번째 인자로 클래스(cls)가 전달됨


```python
class MyClass :
    
    @classmethod
    def class_method(cls, arg1, ...): # cls가 핵심. 클래스 자기자신이 들어감
        
MyClass.class_method(...)
```

* 인스턴스 메서드와의 차이
    * 클래스 내부에 있는 값들을 설정하고 수정
    * 인스턴스 메서드는 인스턴스가 사용할 값들을 설정하고 수정

## 스태틱 메서드
* 클래스가 사용할 메서드
* @staticmethod 데코레이터를 사용하여 정의
* 호출 시, self, cls 둘중 어느 것도 전달되지 않음(클래스 정보에 접근/ 수정 불가)


```python
class MyClass:
    @staticmethod
    def class_method(arg1 ...):
        pass
```

## 메서드 정리


```python
class Korean:
    nation = "ROK"
    code = "KR"
    
    def __init__(self, name, age):
        self.name = name
        self.age = age
        
    def talk(self):
        return f"안녕하세요. {self.name}입니다."
        
    @classmethod
    def info(cls):
        return (cls.nation, cls.code)
    
    def anthem():
        return "동해물과 백두산이 마르고 닳도록..."
```


```python
# 스태틱 메서드
# 인스턴스는 스태틱 메서드 호출 가능
Korean.anthem()
```


    '동해물과 백두산이 마르고 닳도록...'


```python
# 클래스 메서드
# 인스턴스는 클래스 메서드 호출 가능
Korean.info()
```


    ('ROK', 'KR')


```python
# 인스턴스 메서드
kim = Korean("김철수", 100)
kim.talk()
```


    '안녕하세요. 김철수입니다.'

* 스태틱 메서드는 메서드 내부에서 어떠한 객체의 정보도 활용하지 않으며, 클래스가 호출하여 활용함
* 클래스 메서드는 메서드 내부에서 클래스 정보(cls)를 활용, 클래스가 호출하여 활용함.
* 인스턴스 메서드는 메서드 내부에서 인스턴스 정보(self)를 활용, 인스턴스가 호출하여 활용함

* 인스턴스는 나머지 메서드도 접근 가능하나, 일반적으로 사용하지 않는 것이 원칙이다.
* 클래스는 클래스메서드와 스태틱 메서드를 사용한다.
    * 클래스 속성 접근 여부에 따라서 정하면 된다.

# 클래스 변수와 인스턴스 변수

## 클래스 변수와 인스턴스 변수
* 클래스 변수
    * 클래스 정의 안에(+인스턴스 메서드 밖에) 선언
    * 특정 클래스 인스턴스에 묶여 있지 않음
    * 클래스 자체의 내용을 저장
    * 같은 클래스에서 생성된 모든 객체는 동일한 클래스 변수를 공유

* 인스턴스 변수
    * 항상 특정 인스턴스에 묶여 있음
    * 클래스에 저장되지 않고 클래스에서 생성된 개별 객체에 저장
    * 인스턴스 마다 완전히 독립적이므로 변수의 값을 수정하면 오로지 해당 객체에만 영향을 미침.

## 클래스 변수와 인스턴스 변수의 함정
* 새로운 Cat 인스턴스를 만들고 각 인스턴스는 name이라는 인스턴스 변수를 얻음


```python
class Cat:
    num_tails = 1
    
    def __init__(self, name):
        self.name = name
```


```python
alice = Cat("Alice")
james = Cat("james")
```

* Cat 클래스 혹은 각 인스턴스에서 직접 클래스 변수에 접근 다능


```python
print(Cat.num_tails)
print(alice.num_tails)
print(james.num_tails)
```

    1
    1
    1

```python
Cat.num_tails = 2
print(alice.num_tails)
print(james.num_tails)
```

    2
    2


* 클래스 변수를 다시 원래 값으로 변경


```python
Cat.num_tails = 1
```

* 인스턴스로 클래스 변수를 다시 변경해보자


```python
james.num_tails = 2
```

* 그런데 클래스 변수가 변경 된 것이 아닌 james 인스턴스에 num_tails 인스턴스 변수가 새로 생긴 것임
    * 더이상 james는 클래스 변수 num_tais에 접근할 수 없다.


```python
print(alice.num_tails) # 클래스 변수
print(james.num_tails) # 인스턴스 변수
print(Cat.num_tails)   # 클래스 변수
```

    1
    2
    1


### 정리
* 클래스 변수는 모든 클래스 인스턴스에서 공유하는 데이터를 위한 변수
* 인스턴스 변수는 각 인스턴스에 고유한 데이터를 위한 것
* 클래스 변수는 동일한 이름의 인스턴스 변수에 의해 가려질 수 있기 때문에 주의해야함  
(버그나 원치 않은 동작을 유발할 수 있음)

# 객체 비교

## == & is
* ==
    * 동등한(equal)
    * 변수가 참조하는 객체가 동등한(내용이 같은) 경우 True
    * 두 객체가 같아 보이지만 실제로 동일한 대상을 가리키고 있다고 확인해 준 것은 아님
* is
    * 동일한(identical)
    * 두 변수가 동일한 객체를 가리키는 경우 True
    * 메모리 주소 상에서 완전히 같은 주소를 바라보는 경우

# 인스턴스 / 클래스 / 스태틱 메서드

## 인스턴스 / 클래스 / 스태틱 메서드
* 인스턴스 메서드
    * self 매개변수를 통해 동일한 객체에 정의된 속성 및 다른 메서드에 자유롭게 접근 가능
    * 뿐만 아니라 클래스 자체에 접근할 수 있음
        * 즉, 인스턴스 메서드가 클래스 상태를 수정할 수도 있음
* 클래스 메서드
    * 클래스를 가리키는 cls 매개변수를 받음
    * cls인자에만 접근할 수 있기 때문에 객체 인스턴스 상태를 수정할 수 없음
* 스태틱 메서드
    * 임의 개수의 매개 변수를 받을 수 있지만, self나 cls 매개 변수는 사용하지 않음.
    * 즉, 객체 상태나 클래스 상태를 수정할 수 없음
    * 일반 함수처럼 동작하지만, 클래스에 귀속되는 특징


```python
class MyClass:
    def method(self):
        return "instance method", self
    
    @classmethod
    def classmethod(cls):
        return "class method", cls
    
    @staticmethod
    def staticmethod():
        return "static method"
```

* 인스턴스 메서드를 호출한 결과


```python
obj = MyClass()
```


```python
obj.method() ==  MyClass.method(obj) # True
```

* 인스턴스는 클래스 메서드와 스태틱 메서드 모두 접근할 수 있음
    * 그러나 접근해서 변경하려는 행위는 해서는 안됨.


```python
obj.classmethod() # ('class method', __main__.MyClass)
```


```python
MyClass.classmethod() # ('class method', __main__.MyClass)
```


```python
obj.staticmethod() # 'static method'
```

* 클래스 자체에서 각 타입의 메서드를 호출하는 경우
* 인스턴스 메서드는 호출할 수 없음
    * 파이썬이 self 인자를 채울 수 있는 방법이 없으므로 TypeError 예외 발생


```python
MyClass.classmethod() # ('class method', __main__.MyClass)
```


```python
MyClass.staticmethod() # 'static method'
```


```python
MyClass.method()
```


```shell
---------------------------------------------------------------------------
TypeError                                 Traceback (most recent call last)
<ipython-input-95-88f17b4fa117> in <module>
----> 1 MyClass.method()
TypeError: method() missing 1 required positional argument: 'self'
```


## 스태틱 메서드는 언제 사용해야 할까?
* 스태틱 메서드는 self, cls 인자를 취하지 않기 때문에 사용에 있어 큰 제약이 있어 보임
    * 하지만, 반대로 특정한 메서드가 주변의 다른 것들과 독립적일 수 있음을 뜻하는 것이기도 함
* 스태틱 메서드와 클래스 메서드를 사용하는 것은 개발자의 의도를 전달하는 동시에 개발자가 자신의 의도를 강제해 버그러 인해 설계를 깨드리지 않도록 함
* self, cls 인자를 전달하지 않기 때문에 객체 인스턴스, 클래스 상태에 접근할 수 없음을 보장
* 또한 일반함수를 사용하는 것 처럼 실행할 수 있기 때문에 객체 지향 프로그래밍과 절차 지향 프로그래밍 스타일 사이를 연결하는 역할을 하기도 함

## 인스턴스 메서드 / 클래스 메서드 / 정적 메서드 정리
* 인스턴스 메서드는 클래스 인스턴스가 필요하며 self를 통해 인스턴스에 접근
* 클래스 메서드는 클래스 인스턴스가 필요하지 않음
    * 인스턴스(self)에는 접근할 수 없지만 cls를 통해 클래스 자체에 접근할 수 있음
* 스태틱 메서드는 self, cls에 접근할 수 없으며, 일반함수처럼 작동하지만 자신이 정의된 클래스의 이름 공간에 속함
* 스태틱 및 클래스 메서드는 클래스 설계에 대한 개발자의 의도를 전달하고 강제함
    * 이러한 점을 통해 코드의 유지 보수를 하는 데 많은 도움을 줄 수 있음

# 상속

## Why 객체 지향 프로그래밍 (복습)
* 현실 세계를 프로그램 설계에 반영(추상화)
* 그러나 추상화만으로 다양하게 변화하는 현실세계에 대응할 수 있을까?
* Person => Professor, Student

## 상속
* 클래스는 상속이 가능함
    * 모든 파이썬 클래스는 object를 상속 받음
* 상속을 통해 객체 간의 관례를 구축
* 부모 클래스의 속성, 메서드가 자식 클래스에 상속되므로 코드 재사용성이 높아짐


```python
class ChilcClass(ParentClass):
    pass
```

## 상속 - 상속 없이 구현 하는 경우
* 학생 / 교수 정보를 나타내기 어려움


```python
class Person :
    def __init__(self, name, age):
        self.name = name
        self.age = age
        
    def talk(self):
        print(f"반갑습니다. {self.name}입니다.")
```


```python
s1 = Person("김학생", 23)
```


```python
s1.talk() # 반갑습니다. 김학생입니다.
```

```python
p1 = Person("박교수",45)
```


```python
p1.talk() # 반갑습니다. 박교수입니다.
```

```python
s1.gpa = 4.5
```


```python
p1.department = "컴퓨터 공학과"
```

* 그래서 나눠 보았다.
* 메서드 중복 정의


```python
class Professor :
    def __init__(self, name, age, department):
        self.name = name
        self.age = age
        self.department = department
        
    def talk(self):
        print(f"반갑습니다. {self.name}입니다.")
```


```python
class Student :
    def __init__(self, name, age, gpa):
        self.name = name
        self.age = age
        self.gpa = gpa
        
    def talk(self):
        print(f"반갑습니다. {self.name}입니다.")
```


```python
p1 = Professor("박교수", 49, "컴퓨터 공학과")
```


```python
s1 = Student("김학생", 20, 3.5)
```

* 코드 중복이 많이 생김

## 상속
* 상속을 통한 메서드, 재사용


```python
class Person :
    def __init__(self, name, age):
        self.name = name
        self.age = age
        
    def talk(self):
        print(f"반갑습니다. {self.name}입니다.")
        
class Professor(Person) :
    def __init__(self, name, age, department):
        self.name = name
        self.age = age
        self.department = department
        
class Student(Person) :
    def __init__(self, name, age, gpa):
        self.name = name
        self.age = age
        self.gpa = gpa
```


```python
p1 = Professor("박교수", 49, "컴퓨터 공학과")
```


```python
s1 = Student("김학생", 20, 3.5)
```


```python
p1.talk(), s1.talk() # 재사용 가능해짐
```

    반갑습니다. 박교수입니다.
    반갑습니다. 김학생입니다.

## 상속 isinstance
* isinstance(subclass, parentclass) 이면 True


```python
isinstance(p1, Person), isinstance(s1, Person), isinstance(s1, Professor)
```


    (True, True, False)

## issubclass
* is subclass(class, classinfo)


```python
issubclass(bool, int), issubclass(int, float), issubclass(Professor, Person)
```


    (True, False, True)

## super()
* 자식 클래스에서 부모 클래스를 사용하고 싶은 경우
    * super를 통해 반복을 삭제


```python
class Person:
    def __init__( self, name, age, number, email) :
        self.name = name
        self.age = age
        self.nuumber = number
        self.email = email
        
class Student(Person):
    def __init__(self, name, age, number, email, student_id):
        super().__init__(name, age, number, email)
        self.student_id = student_id
```

## 메서드 오버라이딩(method overriding)
* 상속받은 메서드를 재정의
    * 상속 받은 클래스에서 같은 이름의 메서드로 덮어씀
    * 부모 클래스의 메서드를 실행시키고 싶은 경우 super를 활용


```python
class Person:
    def __init__(self, name):
        self.name = name
        
    def talk(self):
        print(f"반갑습니다. {self.name}입니다.")
# 자식 클래스 - Pf
class Professor(Person):
    def talk(self):
        print(f"{self.name}일세.")
# 자식 클래스 - Std
class Student(Person):
    def talk(self):
        super().talk()
        print(f"저는 학생입니다.")
```


```python
pr1 = Person("아무개")
p1 = Professor("김교수")
s1 = Student("이학생")
```


```python
pr1.talk(), p1.talk(), s1.talk()
```

    반갑습니다. 아무개입니다.
    김교수일세.
    반갑습니다. 이학생입니다.
    저는 학생입니다.

## 상속 정리
* 파이썬의 모든 클래스는 object로 부터 상속됨
* 부모 클래스의 모든 요소(속성, 메서드)가 상속됨
* super()를 통해 부모 클래스의 요소를 호출할 수 있음
* 메서드 오버라이딩을 통해 자식 클래스에서도 재정의 할 수 있음
* 상속 관계에서의 이름 공간은 인스턴스, 자식 클래스, 부모 클래스 순으로 탐색

## 다중 상속
* 두개 이상의 클래스를 상속 받는 경우 다중 상속이 됨
    * 상속 받은 모든 클래스의 요소를 활용 가능
    * 중복된 속성이나 메서드가 있는 경우 상속 순서에 의해 결정
* 순서에 신경 쓰면 됨


```python
class A:
    name = 'A'
    def a(self):
        print("A method called")

class B:
    name = 'B'
    def b(self):
        print("B method called")
        
class C(A, B):
    pass

class D(B, A):
    pass
```


```python
c = C() # 상속 순서 A, B 
c.name # 'A'
```


```python
d = D() # 상속 순서 C, D
d.name # 'B'
```


```python
c.a() # A method called
```

```python
c.b() # B method called
```
