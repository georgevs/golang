package main

type Bag[T comparable] struct {
	items map[T]bool
}

func NewBag[T comparable]() *Bag[T] {
	return &Bag[T]{items: make(map[T]bool)}
}

func (b *Bag[T]) Add(x T) {
	b.items[x] = true
}

func (b *Bag[T]) Size() int {
	return len(b.items)
}
