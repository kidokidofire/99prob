package main

import (
	"fmt"
	"strings"
)

// 配列や文字列を指定した位置で分けるsplit関数を実装せよ。
// split([1, 2, 3, 4], 2) // [[1, 2], [3, 4]]
// split('abcdefghik', 3) // ['abc', 'defghik']

func splitInt(arrays []int, start int) [][]int {
	var shifted []int
	for i, v := range arrays {
		if i < start {
			shifted = append(shifted, v)
			arrays = arrays[1:]
		}
	}
	return [][]int{shifted, arrays}
}
func splitString(str string, start int) []string {
	var shifted string
	var strArray = strings.Split(str, "")
	for i, v := range strArray {
		if i < start {
			shifted = shifted + v
			strArray = strArray[1:]
		}
	}
	return []string{shifted, strings.Join(strArray, "")}
}
func main() {
	fmt.Println(splitInt([]int{1, 2, 3, 4}, 2))
	fmt.Println(splitString("abcdefghik", 3))
}
