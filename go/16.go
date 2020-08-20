package main

import (
	"fmt"
	"strings"
)

// 問16: 要素の削除
// nの倍数の位置の要素を配列や文字列から削除するdrop関数を実装せよ。
// drop([1, 2, 3, 4], 2) // [1, 3]
// drop('abcdefghik', 3) // 'abdeghk'

func dropInt(arrays []int, times int) []int {
	var results []int
	for i, v := range arrays {
		if (i+1)%times != 0 {
			results = append(results, v)
		}
	}
	return results
}
func dropString(str string, times int) string {
	var results string
	var strArray = strings.Split(str, "")
	for i, v := range strArray {
		if (i+1)%times != 0 {
			results = results + v
		}
	}
	return results
}
func main() {
	fmt.Println(dropInt([]int{1, 2, 3, 4}, 2))
	fmt.Println(dropString("abcdefghik", 3))
}
