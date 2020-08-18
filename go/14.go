package main

import (
	"fmt"
	"strings"
)

// 問14: 要素の複製
// 配列や文字列の要素を複製するdupli関数を実装せよ。

// dupli([1, 2, 3]) // [1, 1, 2, 2, 3, 3]
// dupli('abc') // aabbcc

func dupliInt(arrays []int) []int {
	var results []int
	for _, v := range arrays {
		results = append(results, v)
		results = append(results, v)
	}
	return results
}
func dupliString(str string) string {
	var results string
	var strArray = strings.Split(str, "")
	for _, v := range strArray {
		results = results + v + v
	}
	return results
}
func main() {
	fmt.Println(dupliInt([]int{1, 2, 3}))
	fmt.Println(dupliString("abc"))
}
