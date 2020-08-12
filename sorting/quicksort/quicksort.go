package quicksort

type QuickSort struct {
	Array []int
}

func (s QuickSort) Sort() []int {
	return s.qSort(0, len(s.Array)-1)
}

func (s QuickSort) qSort(start, end int) []int {
	if start < end {
		pivotIndex := s.partition(start, end)
		s.qSort(start, pivotIndex-1)
		s.qSort(pivotIndex+1, end)
	}
	return s.Array
}

func (s QuickSort) partition(start, end int) int {
	pivot := s.Array[end]
	pivotIndex := start
	i := start
	for ; i < end; i++ {
		if s.Array[i] < pivot {
			s.swap(i, pivotIndex)
			pivotIndex++
		}
	}
	s.swap(end, pivotIndex)

	return pivotIndex
}

func (s QuickSort) swap(i, j int) {
	s.Array[i], s.Array[j] = s.Array[j], s.Array[i]
}
