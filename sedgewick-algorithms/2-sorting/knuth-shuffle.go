package main

import (
	"math/rand"
)

func Shuffle(xs []int) {
	// rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(xs); i++ {
		j := rand.Intn(i + 1)
		xs[i], xs[j] = xs[j], xs[i]
	}
}
