package main
import (
	"fmt"
	"strings"
)

// 配列や文字列から同じ要素の繰り返しを排除して返すcompress関数を実装せよ。
// compress([1, 1, 2, 1, 2, 2, 3, 3, 3, 3]) // [ 1, 2, 1, 2, 3 ]
// compress('aaaabccaadeeee') // abcade

func compressInt(arrays []int) ([]int){
	var results []int
	for _, v := range arrays {
		if (!containsInt(results, v)) {
			results = append(results, v)
		}
	}
	return results
}
func compressString(str string) (string){
	var results []string
	for _, v := range strings.Split(str, "") {
		if (!containsString(results, v)) {
			results = append(results, v)
		}
	}
	return strings.Join(results, "")
}
func containsInt(array []int, element int) bool {
	for _, v := range array {
		if element == v {
			return true
		}
	}
	return false
}
func containsString(array []string, element string) bool {
	for _, v := range array {
		if element == v {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(compressInt([]int{1, 1, 2, 1, 2, 2, 3, 3, 3, 3}))
	fmt.Println(compressString("aaaabccaadeeee"))
}
