package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

// 配列や文字列をランダムに並び替えるrndPermu関数を実装せよ。
// rndPermu('abcdef') // badcef など

func rndPermuInt(array []int) []int {
	result := []int{}
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	for _, index := range rand.Perm(len(array)) {
		result = append(result, array[index])
	}
	return result
}
func rndPermuString(str string) string {
	result := ""
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	for _, index := range rand.Perm(len(str)) {
		result = result + string(str[index])
	}
	return result
}

func main() {
	fmt.Println(rndPermuInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(rndPermuString("abcdef"))
}
