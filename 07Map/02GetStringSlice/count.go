// Author : Github YoonBaek
// 슬라이스로 개표해보기.

package main

import (
	"fmt"
	"log"

	getString "github.com/YoonBaek/learngo/07Map/01GetString"
)

func main() {
	lines, err := getString.GetStrings("../votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 후보자들의 이름을 저장할 변수입니다.
	var names []string
	// 각 이름이 나타나는 횟수를 저장합니다.
	var counts []int

	// 파일의 각 라인을 읽어오며 처리합니다.
	for _, line := range lines {
		matched := false
		// names 슬라이스를 돌며 순회합니다.
		for i, name := range names {
			// 읽어 온 이름이 name 과 일치하면
			if name == line {
				// 현재 name의 인덱스 데로 카운트를 추가합니다.
				counts[i]++
				matched = true
			}
		}
		// 일치하는 이름이 없으면 이름을 추가하고 해당 자리의 count 초깃값을 입력합니다.
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s : %d\n", name, counts[i])
	}
}

/*
	이 방식은 직관적이지만, 이름이 들어올 때 마다 매번 names 슬라이스를 순회하며 검색해야한다는 단점이 있습니다.\
	이는 후보자가 적은 수라면 크게 문제가 되지 않을 수도 있지만,
	스위스와 같이 누구나 피선거자가 될 수 있는 나라의 투표자라면 문제가 될 수 있습니다.
	그리고 매번 슬라이스를 순회하는 것은 효율성 면에서도 좋지 않습니다.

	이 때, 굳이 names 슬라이스를 돌지 않고도 바로 해당 데이터를 찾을 수 있게 ㄷ와주는 자료구조가
	바로 map입니다.
*/
