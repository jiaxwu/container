package deque

import "container/list"

type Stack[T any] struct {
	l *list.List
}

func New[T any]() *Stack[T] {
	return &Stack[T]{l: list.New()}
}

// Push 入栈
func (s *Stack[T]) Push(elem T) {
	s.l.PushBack(elem)
}

// Pop 出栈
func (s *Stack[T]) Pop() T {
	return s.l.Remove(s.l.Back()).(T)
}

// Peek 栈顶元素
func (s *Stack[T]) Peek() T {
	return s.l.Back().Value.(T)
}

// Len 栈元素个数
func (s *Stack[T]) Len() int {
	return s.l.Len()
}

// Empty 栈是否为空
func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}
