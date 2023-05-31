package util

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	mathrand "math/rand"
	"testing"
)

func TestRand(t *testing.T) {

	n := 10  // 10개의 랜덤 숫자 생성
	m := 100 // 랜덤 숫자 범위 (0 ~ 99)

	nums := make(map[int]bool)

	for len(nums) < n {
		num := mathrand.Intn(m)
		if !nums[num] {
			nums[num] = true
		}
	}

	fmt.Println(nums)

	//numbers := make([]int, 0)
	//
	//for i := 0; i < 100; i++ {
	//	// 1부터 100까지 랜덤한 수 생성
	//	num := randInt01(1, 100)
	//
	//	// 중복 체크
	//	for contains(numbers, num) {
	//		num = randInt01(1, 100)
	//	}
	//
	//	numbers = append(numbers, num)
	//}
	//
	//fmt.Println(numbers)

	var numList [100]int64
	for i := 0; i < 100; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			panic(err)
		}
		numList[i] = num.Int64()
		fmt.Print(numList[i], " ")
	}
	str := ``
}
