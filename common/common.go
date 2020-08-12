package common

import (
	"math/rand"
	"time"
)

func GenRandomIntArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(n-1) + 1
	}
	return array
}
