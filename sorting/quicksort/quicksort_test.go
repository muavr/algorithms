package quicksort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/muavr/algorithms/common"
)

func TestQuickSort(t *testing.T) {
	t.Run("Test emty", func(t *testing.T) {
		qs := QuickSort{[]int{}}
		qs.Sort()
	})

	t.Run("Single element", func(t *testing.T) {
		qs := QuickSort{[]int{1}}
		qs.Sort()
	})

	t.Run("Sorted pair", func(t *testing.T) {
		array := []int{1, 2}
		qs := QuickSort{array}
		qs.Sort()
		if !reflect.DeepEqual(array, []int{1, 2}) {
			t.Errorf("Array is not sorted %v", array)
		}
	})

	t.Run("Unsorted pair", func(t *testing.T) {
		array := []int{2, 1}
		qs := QuickSort{array}
		qs.Sort()
		if !reflect.DeepEqual(array, []int{1, 2}) {
			t.Errorf("Array is not sorted %v", array)
		}
	})

	t.Run("Odd number of elements", func(t *testing.T) {
		array := []int{2, 1, 5, 7, 0}
		qs := QuickSort{array}
		qs.Sort()
		if !reflect.DeepEqual(array, []int{0, 1, 2, 5, 7}) {
			t.Errorf("Array is not sorted %v", array)
		}
	})

	t.Run("Even number of elements", func(t *testing.T) {
		array := []int{2, 1, 5, 7, 0, 3}
		qs := QuickSort{array}
		qs.Sort()
		if !reflect.DeepEqual(array, []int{0, 1, 2, 3, 5, 7}) {
			t.Errorf("Array is not sorted %v", array)
		}
	})

	t.Run("Sorted array", func(t *testing.T) {
		array := []int{0, 1, 2, 3, 4, 5}
		qs := QuickSort{array}
		qs.Sort()
		if !reflect.DeepEqual(array, []int{0, 1, 2, 3, 4, 5}) {
			t.Errorf("Array is not sorted %v", array)
		}
	})

	t.Run("Reversed array", func(t *testing.T) {
		array := []int{5, 4, 3, 2, 1, 0}
		qs := QuickSort{array}
		qs.Sort()
		if !reflect.DeepEqual(array, []int{0, 1, 2, 3, 4, 5}) {
			t.Errorf("Array is not sorted %v", array)
		}
	})

	t.Run("Identical values", func(t *testing.T) {
		array := []int{0, 0, 0, 0, 0}
		qs := QuickSort{}
		qs.Sort()
		if !reflect.DeepEqual(array, []int{0, 0, 0, 0, 0}) {
			t.Errorf("Array is not sorted %v", array)
		}
	})

	t.Run("Random values", func(t *testing.T) {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 100; i++ {
			n := rand.Intn(98) + 2
			array := common.GenRandomIntArray(n)
			copiedArray := make([]int, len(array))
			copy(copiedArray, array)
			qs := QuickSort{array}
			qs.Sort()
			sort.Ints(copiedArray)
			if !reflect.DeepEqual(array, copiedArray) {
				t.Errorf("Array is not sorted %v", array)
			}
		}
	})
}
