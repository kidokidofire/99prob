package main

import (
	"fmt"
	"strings"
)

// 問15: 要素の複製２
// 指定された回数だけ配列や文字列の要素を複製するrepli関数を実装せよ。
// repli([1, 2, 3], 3) // [1, 1, 1, 2, 2, 2, 3, 3, 3]
// repli('abc', 3) // aaabbbccc

func repliInt(arrays []int, times int) []int {
	var results []int
	for _, v := range arrays {
		for i := 0; i < times; i++ {
			results = append(results, v)
		}
	}
	return results
}
func repliString(str string, times int) string {
	var results string
	var strArray = strings.Split(str, "")
	for _, v := range strArray {
		for i := 0; i < times; i++ {
			results = results + v
		}
	}
	return results
}
func main() {
	fmt.Println(repliInt([]int{1, 2, 3}, 3))
	fmt.Println(repliString("abc", 3))
}
