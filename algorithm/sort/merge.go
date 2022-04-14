package sort

import "golang.org/x/exp/constraints"

func MergeSort[T constraints.Ordered](arr []T) {
	merge(arr, make([]T, len(arr)), 0, len(arr)-1)
}

func merge[T constraints.Ordered](arr []T, aux []T, lo int, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2

	merge(arr, aux, lo, mid)
	merge(arr, aux, mid+1, hi)
	mergeSort(arr, aux, lo, hi, mid)
}

func mergeSort[T constraints.Ordered](arr []T, aux []T, lo int, hi int, mid int) {
	for i := lo; i <= hi; i++ {
		aux[i] = arr[i]
	}

	lt, rt := lo, mid+1

	for i := lo; i <= hi; i++ {
		if lt > mid {
			arr[i] = aux[rt]
			rt++
		} else if rt > hi {
			arr[i] = aux[lt]
			lt++
		} else if less(aux, lt, rt) {
			arr[i] = aux[lt]
			lt++
		} else {
			arr[i] = aux[rt]
			rt++
		}
	}
}
