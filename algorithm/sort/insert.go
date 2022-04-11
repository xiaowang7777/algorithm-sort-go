package sort

import "golang.org/x/exp/constraints"

func InsertSort[T constraints.Ordered](arr []T) {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i; j < len(arr); j++ {
			if less(arr, j, minIdx) {
				minIdx = j
			}
		}
		if i != minIdx {
			exch(arr, minIdx, i)
		}
	}
}
