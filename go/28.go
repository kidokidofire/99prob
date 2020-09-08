package main

import (
	"fmt"
	"sort"
)

// 配列の配列を、要素の長さでソートするlsort関数を実装せよ。
// また、要素の長さの頻度順にソートするlfsort関数を実装せよ。
// lsort(["abc","de","fgh","de","ijkl","mn","o"]) // ["o","de","de","mn","abc","fgh","ijkl"]
// lfsort(["abc", "de", "fgh", "de", "ijkl", "mn", "o"]) // ["ijkl","o","abc","fgh","de","de","mn"]

func encodeInt(arrays []int) [][]int {
	results := [][]int{}
	currentNum := 0
	currentInt := 0
	for i, v := range arrays {
		if i == 0 || v == arrays[i-1] {
			currentNum = currentNum + 1
			currentInt = v
		} else {
			results = append(results, []int{currentNum, currentInt})
			currentNum = 1
			currentInt = v
		}
	}
	return append(results, []int{currentNum, currentInt})
}

func lsort(str []string) []string {
	sort.Slice(str, func(i, j int) bool { return len(str[i]) < len(str[j]) })
	return str
}
func lfsort(str []string) []string {
	lengthArray := []int{}
	for _, v := range str {
		lengthArray = append(lengthArray, len(v))
	}
	sort.Slice(lengthArray, func(i, j int) bool { return lengthArray[i] < lengthArray[j] })
	// fmt.Println(lengthArray) // [1 2 2 2 3 3 4]

	encodeMap := map[int]int{}
	for _, v := range encodeInt(lengthArray) {
		encodeMap[v[1]] = v[0]
	}
	// fmt.Println(encodeMap) // map[1:1 2:3 3:2 4:1]

	sort.Slice(str, func(i, j int) bool { return encodeMap[len(str[i])] < encodeMap[len(str[j])] })
	return str
}
func main() {
	fmt.Println(lsort([]string{"abc", "de", "fgh", "de", "ijkl", "mn", "o"}))
	fmt.Println(lfsort([]string{"abc", "de", "fgh", "de", "ijkl", "mn", "o"}))
}
