package main
import (
	"fmt"
	"strings"
)
// 配列や文字列のn番目の要素を取り出すelementAt関数を実装せよ。
// ただし、最初の要素は0番目ではなく、1番目と数える。
// elementAt([1, 2, 3], 2) // 2
// elementAt('JavaScript', 5) // S

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func elementAtIntArray(intarr []int, index int)(int) {
	return intarr[min(index - 1, len(intarr) -1)]
}
func elementString(str string, index int)(string) {
	return strings.Split(str, "")[min(index - 1, len(str) -1)]
}
func main() {
	fmt.Println(elementAtIntArray([]int{1, 2, 3}, 2))
	fmt.Println(elementString("JavaScript", 5))
	fmt.Println(elementAtIntArray([]int{1, 2, 3}, 1))
	fmt.Println(elementString("JavaScript", 1))
	fmt.Println(elementAtIntArray([]int{1, 2, 3}, 4))
	fmt.Println(elementString("JavaScript", 11))
}