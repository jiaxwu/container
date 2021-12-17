package heap

import "container/heap"

type Heap[T any] struct {
	h *internalHeap[T]
}

func New[T any](less func(e1 T, e2 T) bool) *Heap[T] {
	return &Heap[T]{h: newInternalHeap[T](less)}
}

func (h *Heap[T]) Len() int {
	return h.h.Len()
}

func (h *Heap[T]) Push(x T) {
	heap.Push(h.h, x)
}

func (h *Heap[T]) Pop() T {
	return heap.Pop(h.h).(T)
}

func (h *Heap[T]) Empty() bool {
	return h.Len() == 0
}

type internalHeap[T any] struct {
	l    []T
	less func(e1 T, e2 T) bool
}

func (h *internalHeap[T]) Len() int {
	return len(h.l)
}

func (h *internalHeap[T]) Less(i, j int) bool {
	return h.less(h.l[i], h.l[j])
}

func (h *internalHeap[T]) Swap(i, j int) {
	h.l[i], h.l[j] = h.l[j], h.l[i]
}

func (h *internalHeap[T]) Push(x any) {
	h.l = append(h.l, x.(T))
}

func (h *internalHeap[T]) Pop() any {
	elem := h.l[len(h.l)-1]
	h.l = h.l[:len(h.l)-1]
	return elem
}

func newInternalHeap[T any](less func(e1 T, e2 T) bool) *internalHeap[T] {
	h := &internalHeap[T]{
		less: less,
	}
	heap.Init(h)
	return h
}
