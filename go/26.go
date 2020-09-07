package main

import (
	"fmt"
	"strings"
)

// m個の要素からn個を選んだ組み合わせを返すcombinations関数を実装せよ。
// combinations('abcdef', 2) // ['ab', 'ac', 'ad', 'ae', 'af', 'bc', 'bd', 'be', 'bf', 'cd', 'ce', 'cf', 'de', 'df', 'ef']

func combinations(str string, num int) []string {
	var strArray = strings.Split(str, "")
	result := make([]string, num)
	return strArrayAppend(0, 0, result, num-1, strArray, []string{})
}
func strArrayAppend(index int, next int, result []string, last int, strArray []string, returnArray []string) []string {
	for j := next; j < len(strArray); j++ {
		result[index] = strArray[j]
		if index == last {
			returnArray = append(returnArray, strings.Join(result, ""))
		} else {
			returnArray = append(returnArray, strArrayAppend(index+1, j+1, result, last, strArray, []string{})...)
		}
	}
	return returnArray
}
func combinationsInt(str []int, num int) [][]int {
	result := make([]int, num)
	return intArrayAppend(0, 0, result, num-1, str, [][]int{})
}
func intArrayAppend(index int, next int, result []int, last int, strArray []int, returnArray [][]int) [][]int{
	for j := next; j < len(strArray); j++ {
		result[index] = strArray[j]
		if index == last {
			item := append([]int{}, result...)
			returnArray = append(returnArray, item)
		} else {
			returnArray = append(returnArray, intArrayAppend(index+1, j+1, result, last, strArray, [][]int{})...)
		}
	}
	return returnArray
}

func main() {
	fmt.Println(combinations("abcdef", 2))
	fmt.Println(combinations("abcdef", 3))
	fmt.Println(combinations("abcdef", 4))
	fmt.Println(combinationsInt([]int{1,2,3,4,5,6}, 2))
}
