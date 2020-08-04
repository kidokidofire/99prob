package main
import (
	"fmt"
	"strings"
)
// 配列や文字列を逆順にして返すreverse関数を実装せよ。
// ただし、Array.reverseは使用しないこと。
// また、絵文字の長さは1と数えるようにすること。

// reverse('A man, a plan, a canal, panama!💃') // 💃!amanap ,lanac a ,nalp a ,nam A
// reverse([1, 2, 3, 4]) // [4, 3, 2, 1]

func reverseIntArr(intarr []int)([]int) {
	returnArray := []int{}
	for _, v := range intarr {
		returnArray= append([]int{v}, returnArray...)
	}
	return returnArray
}	
func reverseStr(str string)(string) {
	returnArray := ""
	for _, v := range strings.Split(str, "") {
		returnArray = v + returnArray
	}
	return returnArray
}	
func main() {
	fmt.Println(reverseIntArr([]int{1, 2, 3, 4}))
	fmt.Println(reverseIntArr([]int{123, 456, 789}))
	fmt.Println(reverseStr("A man, a plan, a canal, panama!💃"))
	fmt.Println(reverseStr("💃Hello, World!💃"))
}
