# 1.
num = int(input("숫자 입력 :"))
if num % 2 : # num % 2 == 1
    print("홀수입니다.")
else :
    print("짝수입니다. ")

# 2. 중첩조건문 실습 문제
dust = 500
if dust > 150 :
    print("매우 나쁨")
    if dust > 300 :
        print("실외활동을 자제하세요")
elif dust > 80 :
    print("나쁨")
elif dust > 30 :
    print("보통")
else :
    if dust >= 0 :
        print("좋음")
    else :
        print("wrong value")

# 3. 
input_num = int(input())
a = 1; total = 0
while a <= 5 :
    total += a
    a += 1
print(total)

# 4. 
inputs = input()
for i in inputs :
    print(i)