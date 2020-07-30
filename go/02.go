package main
import (
	"fmt"
	"strings"
)
// 配列や文字列の最後の一つ前の要素を取り出すbutLast関数
// butLast([1, 2, 3, 4]) // 3
// butLast('abcdefghijklmnopqrstuvwxyz') // y
func butLast(arg interface{})(interface{}) {
	str, strok := arg.(string) 
	if (strok) {
		return strings.Split(str, "")[len(str)-2]
	}
	intarr, intok := arg.([]int) 
	if (intok) {
		return intarr[len(intarr)-2]
	}
	return "err";
}
func main() {
	fmt.Println(butLast([]int{1, 2, 3, 4}))
	fmt.Println(butLast("abcdefghijklmnopqrstuvwxyz"))
}