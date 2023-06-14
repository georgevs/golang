package main

import (
	"fmt"
)

type PriorityQueue struct {
	xs []int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{xs: []int{}}
}

func NewPriorityQueueFromList(xs []int) *PriorityQueue {
	h := NewPriorityQueue()
	for _, x := range xs {
		h.Enqueue(x)
	}
	return h
}

func (h *PriorityQueue) Enqueue(x int) {
	h.xs = append(h.xs, x)
	h.bubble(len(h.xs) - 1)
}

func (h *PriorityQueue) Dequeue() int {
	h.swap(0, len(h.xs)-1)
	x := h.xs[len(h.xs)-1]
	h.xs = h.xs[:len(h.xs)-1]
	h.sink(0)
	return x
}

func (h *PriorityQueue) Size() int {
	return len(h.xs)
}

func (h *PriorityQueue) sink(i int) int {
	k := i
	l := h.left(i)
	if l < len(h.xs) && h.xs[l] < h.xs[k] {
		k = l
	}
	r := h.right(i)
	if r < len(h.xs) && h.xs[r] < h.xs[k] {
		k = r
	}
	if k != i {
		h.swap(k, i)
		return h.sink(k)
	}
	return i
}

func (h *PriorityQueue) bubble(i int) int {
	p := h.parent(i)
	if h.xs[i] < h.xs[p] {
		h.swap(i, p)
		return h.bubble(p)
	}
	return i
}

func (h *PriorityQueue) swap(i, j int) {
	h.xs[i], h.xs[j] = h.xs[j], h.xs[i]
}

func (h *PriorityQueue) parent(i int) int {
	return (i - 1) / 2
}

func (h *PriorityQueue) left(i int) int {
	return 2*i + 1
}

func (h *PriorityQueue) right(i int) int {
	return 2*i + 2
}

func (h *PriorityQueue) String() string {
	return fmt.Sprintf("PriorityQueue{%v}", h.xs)
}
