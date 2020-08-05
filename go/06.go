package main
import (
	"fmt"
	"strings"
)

// 配列や文字列が回文かどうかを返すisPalindrome関数を実装せよ。

// isPalindrome([1, 2, 3]) // false
// isPalindrome('たけやぶやけた') // true
// isPalindrome([1, 2, 4, 8, 16, 8, 4, 2, 1]) // true

func isPalindromeIntArr(intarr []int)(bool) {
	for i, _ := range intarr {
		if (i >= len(intarr) - i) {
			return true
		}
		if (intarr[i] != intarr[len(intarr) - 1 - i]) {
			return false
		}
	}
	return false
}	
func isPalindromeStr(str string)(bool) {
	splitedStr := strings.Split(str, "")
	for i, _ := range splitedStr {
		if (i >= len(splitedStr) - i) {
			return true
		}
		if (splitedStr[i] != splitedStr[len(splitedStr) - 1 - i]) {
			return false
		}
	}
	return false
}	
func main() {
	fmt.Println(isPalindromeIntArr([]int{1, 2, 3, 4}))
	fmt.Println(isPalindromeIntArr([]int{1, 2, 4, 8, 16, 8, 4, 2, 1}))
	fmt.Println(isPalindromeStr("たけやぶやけた"))
}