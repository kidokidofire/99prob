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
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	startIndex := rand.Intn(len(array) - selectNum)
	return array[startIndex : startIndex+selectNum]
}
func rndSelectString(str string, selectNum int) string {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	startIndex := rand.Intn(len(str) - selectNum)
	return str[startIndex : startIndex+selectNum]
}

func main() {
	fmt.Println(rndSelectInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 2))
	fmt.Println(rndSelectString("abcdefgh", 3))
}
