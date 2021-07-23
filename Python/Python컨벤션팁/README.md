# Python 컨벤션 팁

## convention

1. 리스트 / 배열을 담는 변수명은 복수형으로
2. 함수 이름은 동사형으로 짓는다. => 이름에서 return 타입을 추측할 수 있으면 더 좋다.

```python
def is_adult(age) :
    if age > 20 :
        return True
    else :
        return False
def get_lotto_numbers() :
    import random
    numbers = random.sample(range(1, 46), 6)
    return numbers
```

