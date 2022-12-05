package utils

type Stack []rune

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str rune) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 'a', false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) PopR() (rune, bool) {
	if s.IsEmpty() {
		return 'a', false
	} else {
		element := (*s)[0] // Index into the slice and obtain the element.
		*s = (*s)[1:]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Peek() (rune, bool) {
	if s.IsEmpty() {
		return 'a', false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}

}
func (s *Stack) Reverse() {

	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
