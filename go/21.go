package main

import (
	"fmt"
	"strings"
)

// 配列や文字列の指定した位置に要素を挿入するinsertAt関数を実装せよ。

// insertAt(5, [1, 2, 3, 4], 3) // [1, 2, 5, 3, 4]
// insertAt('X', 'abcd', 2) // aXbcd

func insertAtInt(add int, array []int, index int) []int {
	lastHalf := append([]int{add}, array[index-1:]...)
	return append(array[:index-1], lastHalf...)
}
func insertAtString(add string, str string, index int) string {
	strArray := strings.Split(str, "")
	return strings.Join(strArray[:index-1], "") + add + strings.Join(strArray[index-1:], "")
}

func main() {
	fmt.Println(insertAtInt(5, []int{1, 2, 3, 4}, 3))
	fmt.Println(insertAtString("X", "abcd", 2))
}
