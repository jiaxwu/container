package deque

import "container/list"

type Deque[T any] struct {
	l *list.List
}

func New[T any]() *Deque[T] {
	return &Deque[T]{l: list.New()}
}

// AddFirst 从队头入队
func (d *Deque[T]) AddFirst(elem T) {
	d.l.PushFront(elem)
}

// AddLast 从队尾入队
func (d *Deque[T]) AddLast(elem T) {
	d.l.PushBack(elem)
}

// RemoveFirst 从队头出队
func (d *Deque[T]) RemoveFirst() T {
	return d.l.Remove(d.l.Front()).(T)
}

// RemoveLast 从队尾出队
func (d *Deque[T]) RemoveLast() T {
	return d.l.Remove(d.l.Back()).(T)
}

// Len 队列元素个数
func (d *Deque[T]) Len() int {
	return d.l.Len()
}

// Empty 队列是否为空
func (d *Deque[T]) Empty() bool {
	return d.Len() == 0
}
