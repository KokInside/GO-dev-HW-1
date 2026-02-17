package stack

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

func (s *Stack[T]) Top() T {
	if len(s.stack) != 0 {
		return s.stack[len(s.stack)-1]
	}
	panic("Top empty stack")
}

func (s *Stack[T]) Size() int {
	return len(s.stack)
}
