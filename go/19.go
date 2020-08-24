package main

import (
	"fmt"
	"strings"
)

// 配列や文字列の要素を左にn個ずらすrotate関数を実装せよ。
// 負の数を渡したときは右にずらすようにすること。
// rotate([1, 2, 3], 1) // [2, 3, 1]
// rotate('abcdefgh', 3) // defghabc
// rotate('abcdefgh', -2) // ghabcdef

func rotateInt(arrays []int, shift int) []int {
	var shifted []int
	var shiftLength int
	if shift < 0 {
		shiftLength = shift + len(arrays)
	} else {
		shiftLength = shift
	}
	for i, v := range arrays {
		if i < shiftLength {
			shifted = append(shifted, v)
			arrays = arrays[1:]
		}
	}
	return append(arrays, shifted...)
}
func rotateString(str string, shift int) string {
	var shifted string
	var shiftLength int
	var strArray = strings.Split(str, "")
	if shift < 0 {
		shiftLength = shift + len(strArray)
	} else {
		shiftLength = shift
	}
	for i, v := range strArray {
		if i < shiftLength {
			shifted = shifted + v
			strArray = strArray[1:]
		}
	}
	return strings.Join(strArray, "") + shifted
}
func main() {
	fmt.Println(rotateInt([]int{1, 2, 3}, 1))
	fmt.Println(rotateInt([]int{1, 2, 3}, -1))
	fmt.Println(rotateString("abcdefgh", 3))
	fmt.Println(rotateString("abcdefgh", -2))
}
