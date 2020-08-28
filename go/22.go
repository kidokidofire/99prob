package main

import (
	"fmt"
)

// 指定された範囲内のすべての整数または文字を含む配列を生成するrange関数を実装せよ。
// range(4, 9) // [4, 5 ,6, 7, 8, 9]
// range('a', 'z') // ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']

func rangeInt(start int, end int) []int {
	returnArray := []int{}
	for i := start; i <= end; i++ {
		returnArray = append(returnArray, i)
	}
	return returnArray
}
func rangeString(start rune, end rune) []string {
	returnArray := []string{}
	for i := int(start); i <= int(end); i++ {
		returnArray = append(returnArray, string(rune(i)))
	}
	return returnArray
}

func main() {
	fmt.Println(rangeInt(4, 9))
	fmt.Println(rangeString('a', 'z'))
}
