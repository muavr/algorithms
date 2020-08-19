package bubblesort

type BubbleSort struct {
	Array []int
}

func (s BubbleSort) Sort() []int {
	for i := 0; i < len(s.Array)-1; i++ {
		swapped := false
		min := i
		for j := 0; j < len(s.Array)-1-i; j++ {
			if s.Array[j] > s.Array[j+1] {
				s.swap(j, j+1)
				swapped = true
				if s.Array[i] < s.Array[min] {
					min = i
				}
			}
		}
		if !swapped {
			break
		}
		s.swap(i, min)
	}

	return s.Array
}

func (s BubbleSort) swap(i, j int) {
	s.Array[i], s.Array[j] = s.Array[j], s.Array[i]
}
