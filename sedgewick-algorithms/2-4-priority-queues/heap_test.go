/*
	Build:
		go test -v
*/

package main

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	fmt.Println(HeapSort([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(HeapSort([]int{8, 7, 6, 5, 4, 3, 2, 1}))
	fmt.Println(HeapSort([]int{1, 3, 5, 7, 2, 4, 6, 8}))
}
