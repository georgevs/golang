package main

import "golang.org/x/exp/constraints"

func HeapSort[T constraints.Ordered](xs []T) []T {
	Heapify(xs)
	sortHeap(xs)
	return xs
}

func sortHeap[T constraints.Ordered](xs []T) {
	for n := len(xs); n > 1; {
		n--
		swap(xs, 0, n)
		sink(xs, 0, n)
	}
}

func Heapify[T constraints.Ordered](xs []T) {
	n := len(xs)
	for k := n/2 - 1; k >= 0; k-- {
		sink(xs, k, n)
	}
}

// func HeapifyNaive[T constraints.Ordered](xs []T) {
// 	n := len(xs)
// 	for k := 1; k < n; k++ {
// 		bubble(xs, k)
// 	}
// }

// func bubble[T constraints.Ordered](xs []T, i int) {
// 	for p := (i - 1) / 2; less(xs, i, p); i, p = p, (p-1)/2 {
// 		swap(xs, i, p)
// 	}
// }

func sink[T constraints.Ordered](xs []T, i, n int) {
	for k := i; ; i = k {
		l := i*2 + 1
		if l < n && less(xs, l, k) {
			k = l
		}
		r := i*2 + 2
		if r < n && less(xs, r, k) {
			k = r
		}
		if k == i {
			break
		}
		swap(xs, i, k)
	}
}

func swap[T constraints.Ordered](xs []T, i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}

func less[T constraints.Ordered](xs []T, i, j int) bool {
	return xs[i] < xs[j]
}
