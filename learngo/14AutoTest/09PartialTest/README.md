# 부분테스트 실행하는 법
cmd line flag를 활용해 부분 실행 가능
codeMagnet/compare를 기준으로 한 활용법

테스트 코드
```
go test github.com/YoonBaek/learngo/14AutoTest/08CodeMagnet/compare 
```
# verbose flag
```
github.com/YoonBaek/learngo/14AutoTest/08CodeMagnet/compare -v
```
결과
```
=== RUN   TestFirstLarger
    larger_test.go:14: Larger(2, 1) = 1, want 2
--- FAIL: TestFirstLarger (0.00s)
=== RUN   TestSecondLarger
    larger_test.go:22: Larger(4, 8) = 4, want 8
--- FAIL: TestSecondLarger (0.00s)
FAIL
FAIL    github.com/YoonBaek/learngo/14AutoTest/08CodeMagnet/compare     0.004s
FAIL
```
fail이 나오는게 정상입니다. 일부러 fail 결과를 보려고 만든 코드이고, 위에서 볼건 -v 플래그를 적용하면 테스트함수의 이름과 결과를 같이 보여줍니다.
# -run flag
-run 뒤의 키워드가 함수 이름에 들어가 있는 함수만 Test합니다. 부분 테스트를 실행해야할 때 사용합니다.
```
github.com/YoonBaek/learngo/14AutoTest/08CodeMagnet/compare -run First
```
결과
```
--- FAIL: TestFirstLarger (0.00s)
    larger_test.go:14: Larger(2, 1) = 1, want 2
FAIL
FAIL    github.com/YoonBaek/learngo/14AutoTest/08CodeMagnet/compare     0.004s
FAIL
```
