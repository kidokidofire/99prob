package main
import (
	"fmt"
	"strings"
)

// 配列や文字列の同じ要素の繰り返しを配列としてまとめて返すpack関数を実装せよ。
// pack([1, 1, 2, 1, 2, 2, 3, 3, 3, 3]) // [[1, 1], [2], [1], [2, 2], [3, 3, 3, 3]]
// pack('aaaabccaadeeee') // ['aaaa', 'b', 'cc', 'aa', 'd', 'eeee']

func packInt(arrays []int) ([]interface{}){
	var results []interface{}
	var sameInt []int
	for i, v := range arrays {
		if (i == 0 || v == arrays[i - 1]) {
			sameInt = append(sameInt, v)
		} else {
			results = append(results, sameInt)
			sameInt  = []int{v}
		}
	}
	results = append(results, sameInt)
	return results
}
func packString(str string) ([]interface{}){
	var results []interface{}
	var sameChar []string
	var strArray = strings.Split(str, "")
	for i, v := range strArray  {
		if (i == 0 || v == strArray[i - 1]) {
			sameChar= append(sameChar, v)
		} else {
			results = append(results, sameChar)
			sameChar= []string{v}
		}
	}
	return results
}
func main() {
	fmt.Println(packInt([]int{1, 1, 2, 1, 2, 2, 3, 3, 3, 3}))
	fmt.Println(packString("aaaabccaadeeee"))
}
