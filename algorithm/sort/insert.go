package sort

import "golang.org/x/exp/constraints"

func InsertSort[T constraints.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if less(arr, j, j-1) {
				exch(arr, j, j-1)
			} else {
				break
			}
		}
	}
}
