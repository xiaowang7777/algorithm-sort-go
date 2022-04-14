package sort

import "golang.org/x/exp/constraints"

func QuickSort[T constraints.Ordered](arr []T) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort[T constraints.Ordered](arr []T, lo int, hi int) {
	if lo >= hi {
		return
	}
	p := partition(arr, lo, hi)
	quickSort(arr, lo, p-1)
	quickSort(arr, p+1, hi)
}

func partition[T constraints.Ordered](arr []T, lo int, hi int) int {
	i, j := lo, hi

	p := lo

	for true {
		for less(arr, i, p) {
			if i == p || i >= hi {
				break
			}
			i++
		}

		for less(arr, p, j) {
			if p == j || j <= lo {
				break
			}
			j--
		}

		if i >= j {
			break
		}

		exch(arr, i, j)
		i++
		j--
	}
	exch(arr, p, j)

	return j
}
