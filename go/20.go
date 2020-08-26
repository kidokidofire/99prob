package main

import (
	"fmt"
	"strings"
)

// 配列や文字列のn番目の要素を削除するremoveAt関数を実装せよ。
// 削除した要素と処理後の配列を配列に格納し、呼び出した結果として返しなさい。

// removeAt(3, [1, 2, 3]) // [3, [1, 2]]
// removeAt(2, 'abcd') // ['b', 'acd']

func removeAtInt(remove int, array []int) (int, []int) {
	if len(array) < remove {
		return 0, nil
	}
	var removed = array[remove-1]
	var rest = append(array[:remove-1], array[remove:]...)
	return removed, rest
}
func removeAtString(remove int, str string) (string, string) {
	if len(str) < remove {
		return "", ""
	}
	var strArray = strings.Split(str, "")
	var removed = strArray[remove-1]
	var rest = append(strArray[:remove-1], strArray[remove:]...)
	return removed, strings.Join(rest, "")
}

func main() {
	fmt.Println(removeAtInt(3, []int{1, 2, 3}))
	fmt.Println(removeAtString(2, "abcd"))
}
