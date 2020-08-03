package main
import (
	"fmt"
	"strings"
)
// 配列や文字列の長さを返すlength関数を実装せよ。
// ただし、Array.lengthやString.lengthは使用しないこと。
// また、絵文字の長さは1と数えるようにすること。
// length([123, 456, 789]) // 3
// length('💃Hello, World!💃') // 15

func lengthIntArr(intarr []int)(int) {
	lengthMax := 0
	for i, _ := range intarr {
		lengthMax = i
	}
	return lengthMax + 1
}	
func lengthStr(str string)(int) {
	lengthMax := 0
	for i, _ := range strings.Split(str, "") {
		lengthMax = i
	}
	return lengthMax + 1
}	
func main() {
	fmt.Println(lengthIntArr([]int{123, 456, 789}))
	fmt.Println(lengthStr("💃Hello, World!💃"))
}
