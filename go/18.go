package main

import (
	"fmt"
	"strings"
)

// 選択した範囲を配列や文字列から取り出すslice関数を実装せよ。
// ただし、Array.slice、Array.spliceは使用しないこと。
// slice([1, 2, 3, 4], 2, 4) // [2, 3, 4]
// slice('abcdefghik', 3, 7) // cdefg

func sliceInt(arrays []int, start int, end int) []int {
	var resultArray []int
	for i, v := range arrays {
		if start <= i+1 && i+1 <= end {
			resultArray = append(resultArray, v)
		}
	}
	return resultArray
}
func sliceString(str string, start int, end int) string {
	var resultArray = ""
	var strArray = strings.Split(str, "")
	for i, v := range strArray {
		if start <= i+1 && i+1 <= end {
			resultArray = resultArray + v
		}
	}
	return resultArray
}
func main() {
	fmt.Println(sliceInt([]int{1, 2, 3, 4}, 2, 4))
	fmt.Println(sliceString("abcdefghik", 3, 7))
}
