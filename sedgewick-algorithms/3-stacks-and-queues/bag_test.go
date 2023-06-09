package main

import "testing"

func TestBag(t *testing.T) {
	b := NewBag[int]()
	b.Add(1)
	b.Add(2)
	b.Add(2)
	b.Add(3)
	if b.Size() != 3 {
		t.Error()
	}
}
