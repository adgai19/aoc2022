package utils

type Queue[T any] []T

// IsEmpty: check if stack is empty
func (s *Queue[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Queue[T]) Enque(item T) {
	*s = append(*s, item) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Queue[T]) Deque() (T, bool) {
	if s.IsEmpty() {
		var res T
		return res, false
	} else {
		element := (*s)[0] // Index into the slice and obtain the element.
		*s = (*s)[1:]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// func (s *Stack[T]) Peek() (T, bool) {
// 	if s.IsEmpty() {
// 		var res T
// 		return res, false
// 	} else {
// 		index := len(*s) - 1   // Get the index of the top most element.
// 		element := (*s)[index] // Index into the slice and obtain the element.
// 		return element, true
// 	}
//
// }
// func (s *Stack[T]) Reverse() {
//
// 	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
// 		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
// 	}
// }
