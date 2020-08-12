package main

import (
	"fmt"

	"github.com/muavr/algorithms/common"

	"github.com/muavr/algorithms/sorting/quicksort"
)

func main() {
	n := 20
	randomArray := common.GenRandomIntArray(n)
	fmt.Println("random: ", randomArray)
	q := quicksort.QuickSort{Array: randomArray}
	fmt.Println("sorted: ", q.Sort())
}
