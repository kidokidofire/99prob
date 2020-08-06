package main
import (
	"fmt"
)

// ネストしている配列を平坦（一次元配列）にして返すflatten関数を実装せよ。
// ただし、Array.flatは使用しないこと。

// flatten([1, [2, [3, 4], 5]]) // [1, 2, 3, 4, 5]

func flatten(arrays []interface{}) []int{
	var result []int
	for _, s := range arrays {
		switch i := s.(type) {
		case int:
			result = append(result, i)
		case []interface{}:
			result = append(result, flatten(i)...)
		}
	}
	return result
}
func main() {
	fmt.Println(flatten([]interface{}{[]interface{}{1, []interface{}{2, []interface{}{3, 4}, 5}}}))
}
