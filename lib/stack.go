package lib

type Stack[T any] struct {
	Stack []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

// Returns length of the stack.
func (s *Stack[T]) Len() int {
	return len(s.Stack)
}

// Reports if stack is empty (no elements, length is 0).
func (s *Stack[T]) IsEmpty() bool {
	return len(s.Stack) == 0
}

// Peeks at the top element of the stack.
func (s *Stack[T]) Top() T {
	return s.Stack[len(s.Stack)-1]
}

// Appends elements to the top of the stack (last element becomes top).
func (s *Stack[T]) Append(e ...T) {
	s.Stack = append(s.Stack, e...)
}

// Pops and returns top element from the stack.
func (s *Stack[T]) Pop() T {
	top := s.Top()
	s.Stack = s.Stack[:len(s.Stack)-1]
	return top
}
