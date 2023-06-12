/*
	Build:
		go test ./sedgewick-algorithms/2-sorting
*/

package main

import (
	"fmt"
	"testing"
)

func TestKnuthShuffle(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Shuffle(xs)
	fmt.Println(xs)
}
