package sort

import "golang.org/x/exp/constraints"

func ShellSort[T constraints.Ordered](arr []T, step int) {
	for s := step; s > 0; s-- {
		for i := s; i < len(arr); i++ {
			for j := i; j >= s; j -= s {
				if less(arr, j, j-s) {
					exch(arr, j, j-s)
				} else {
					break
				}
			}
		}
	}
}
