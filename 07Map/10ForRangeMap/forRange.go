// Author : Github Yoon Baek
// 앞선 투표 결과를 예쁘게 만들기 위해 결과를 for ...range문으로 풀어줄 겁니다.
// 주의할 점! map은 정렬되지 않은 커렉션이기 때문에 루프가 맵을 무작위 순서로 선회합니다.
// 따라서 for range문을 mao애 사용하는 경우 값이 어떤 순서로 반환될지 알 수 없습니다.
// 따라서 순서를 지정해주는 추가작업이 필요합니다.

package main

import (
	"fmt"
	"log"
	"sort"

	getString "github.com/YoonBaek/learngo/07Map/01GetString"
)

func main() {
	lines, err := getString.GetStrings("../votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	// 아래 코드는 매번 프로그램을 실행할 때마다 순서가 바뀝니다.
	// 순서가 중요한 선거에서 결과가 이렇게 나온다면 좋지 않겠죠.
	for name, count := range counts {
		fmt.Printf("%s gets %d votes.\n", name, count)
	}

	// 그래서 아래와 같이 이름을 정렬해 줍니다.
	names := []string{}
	// 키만 필요한 경우에는 값을 생략해줄 수 있습니다.
	// 값은 키를 따로 언더바로 무시해줘야 합니다.
	for name := range counts {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s gets %d votes.\n", name, counts[name])
	}
}

/*
드디어 개표 프로그램이 완성되었습니다.
배열과 슬라이스만 사용했을 때에는 값을 찾기 위한 많은 추가코드와 연산 비용이 필요했습니다.
하지만 맵을 사용하면 이 모든 과정을 쉽고 비용효율적으로 처리할 수 있습니다.
언제든 컬렉셩네서 값을 검색해야 하는 일이 생긴다면 맵을 사용해보는 것도 한 번 고려해 보세요!
*/
