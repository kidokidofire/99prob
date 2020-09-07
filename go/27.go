package main

import (
	"fmt"
)

// 配列の要素を互いに素な配列にグループ化して返すgroup関数を実装せよ。
// group(["aldo","beat","carla","david","evi","flip","gary","hugo","ida"], 2, 3, 4)
// // [[["aldo","beat"],["carla","david","evi"],["flip","gary","hugo","ida"]],...] 1260個の解
// group(["aldo","beat","carla","david","evi","flip","gary","hugo","ida"], 2, 2, 5)
// // [[["aldo","beat"],["carla","david"],["evi","flip","gary","hugo","ida"]],...] 756個の解

func combinations(str []string, num int) [][]string {
	result := make([]string, num)
	return strArrayAppend(0, 0, result, num-1, str, [][]string{})
}
func strArrayAppend(index int, next int, result []string, last int, strArray []string, returnArray [][]string) [][]string {
	for j := next; j < len(strArray); j++ {
		result[index] = strArray[j]
		if index == last {
			item := append([]string{}, result...)
			returnArray = append(returnArray, item)
		} else {
			returnArray = append(returnArray, strArrayAppend(index+1, j+1, result, last, strArray, [][]string{})...)
		}
	}
	return returnArray
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func groupInt(strArray []string, numArray []int) [][][]string {
	if len(numArray) == 0 {
		return [][][]string{}
	}

	returnArray := [][][]string{}
	numShift := numArray[0]
	numArray = numArray[1:]

	for _, v := range combinations(strArray, numShift) {
		// strArray が [v][leftArray] に分かれる
		leftArray := []string{}
		for _, strItem := range strArray {
			if !contains(v, strItem) {
				leftArray = append(leftArray, strItem)
			}
		}

		// 結果を詰める
		if len(leftArray) == 0 {
			returnArray = append(returnArray, [][]string{v})
		} else {
			for _, left := range groupInt(leftArray, numArray) {
				item := append([][]string{v}, left...)
				returnArray = append(returnArray, item)
			}
		}
	}
	return returnArray
}

func group(strArray []string, numArray []int) [][][]string {
	result := groupInt(strArray, numArray)
	fmt.Println(len(result))
	return result
}

func main() {
	fmt.Println(group([]string{"aldo", "beat", "carla", "david", "evi", "flip", "gary", "hugo", "ida"}, []int{2, 3, 4}))
	fmt.Println(group([]string{"aldo", "beat", "carla", "david", "evi", "flip", "gary", "hugo", "ida"}, []int{2, 2, 5}))
}
