package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

// 長さnの1以上m以下の乱数列を生成するdiffSelect関数を実装せよ。
// diffSelect(6, 49) // [23, 1, 17, 33, 21, 37] など

func diffSelect(number int, rangeNum int) []int {
	if number == 0 {
		return []int{}
	}
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	return append([]int{rand.Intn(rangeNum) + 1}, diffSelect(number-1, rangeNum)...)
}

func main() {
	fmt.Println(diffSelect(6, 49))
}
