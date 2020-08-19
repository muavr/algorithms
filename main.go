package main

import (
	"fmt"

	"github.com/muavr/algorithms/common"

	"github.com/muavr/algorithms/sorting"
	"github.com/muavr/algorithms/sorting/bubblesort"
	"github.com/muavr/algorithms/sorting/quicksort"
)

func main() {
	n := 20
	randomArray := common.GenRandomIntArray(n)
	fmt.Println("random: ", randomArray)
	var q sorting.Sortable
	q = quicksort.QuickSort{Array: randomArray}
	fmt.Println("sorted: ", q.Sort())

	randomArray = common.GenRandomIntArray(n)
	fmt.Println("random: ", randomArray)
	var b sorting.Sortable
	b = bubblesort.BubbleSort{Array: randomArray}
	fmt.Println("sorted: ", b.Sort())
}
