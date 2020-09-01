package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

// 配列や文字列から指定された数ぶんだけランダムに要素を取り出すrndSelect関数を実装せよ。
// rndSelect('abcdefgh', 3) // eda など

func rndSelectInt(array []int, selectNum int) []int {
	if selectNum == 0 {
		return []int{}
	}
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	index := rand.Intn(len(array) - selectNum)
	return append([]int{array[index]}, rndSelectInt(array, selectNum-1)...)
}
func rndSelectString(str string, selectNum int) string {
	if selectNum == 0 {
		return ""
	}
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	index := rand.Intn(len(str) - selectNum)
	return str[index:index+1] + rndSelectString(str, selectNum-1)
}

func main() {
	fmt.Println(rndSelectInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 2))
	fmt.Println(rndSelectString("abcdefgh", 3))
}
