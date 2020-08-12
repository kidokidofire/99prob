package main
import (
	"fmt"
	"strings"
	"strconv"
)

// pack関数を用いて、配列や文字列をランレングス圧縮するencode関数を実装せよ。
// encode([1, 1, 2, 1, 2, 2, 3, 3, 3, 3]) // [[2, 1], [1, 2], [1, 1], [2, 2], [4, 3]]
// encode('aaaabccaadeeee')
// // [[4, 'a'], [1, 'b'], [2, 'c'], [2, 'a'], [1, 'd'], [4, 'e']]

func encodeInt(arrays []int) ([]interface{}){
	var results []interface{}
	var currentNum int = 0
	var currentInt int
	for i, v := range arrays {
		if (i == 0 || v == arrays[i - 1]) {
			currentNum = currentNum + 1
			currentInt = v
		} else {
			results = append(results, []int{currentNum, currentInt})
			currentNum = 1
			currentInt = v
		}
	}
	results = append(results, []int{currentNum, currentInt})
	return results
}
func encodeString(str string) ([]interface{}){
	var results []interface{}
	var currentNum int = 0
	var currentChar string
	var strArray = strings.Split(str, "")
	for i, v := range strArray {
		if (i == 0 || v == strArray[i - 1]) {
			currentNum = currentNum + 1
			currentChar = v
		} else {
			results = append(results, []string{strconv.Itoa(currentNum), currentChar})
			currentNum = 1
			currentChar = v
		}
	}
	results = append(results, []string{strconv.Itoa(currentNum), currentChar})
	return results
}
func main() {
	fmt.Println(encodeInt([]int{1, 1, 2, 1, 2, 2, 3, 3, 3, 3}))
	fmt.Println(encodeString("aaaabccaadeeee"))
}
