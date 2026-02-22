package stack

import "errors"

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) Push(value T) {
	s.stack = append(s.stack, value)
}

func (s *Stack[T]) Pop() {
	if len(s.stack) != 0 {
		s.stack = s.stack[:len(s.stack)-1]
		return
	}
	panic("pop empty stack")
}

func (s *Stack[T]) Top() (T, error) {
	if len(s.stack) != 0 {
		return s.stack[len(s.stack)-1], nil
	}

	var zero T

	return zero, errors.New("Top empty stack")
}

func (s *Stack[T]) Size() int {
	return len(s.stack)
}
