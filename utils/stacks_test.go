package utils

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := make(Stack[int], 0)
	t.Run("Add items to stack", func(t *testing.T) {
		s.Push(5)
		want := Stack[int]{5}
		assert(t, s, want)
	})
	t.Run("check if non-empty list is empty", func(t *testing.T) {
		assert(t, s.IsEmpty(), false)
	})
	t.Run("Pop items", func(t *testing.T) {
		item, err := s.Pop()
		assert(t, item, 5)
		assert(t, err, true)
	})
	t.Run("check if list is empty", func(t *testing.T) {
		assert(t, s.IsEmpty(), true)
	})

	t.Run("Pop items from empty list", func(t *testing.T) {
		item, err := s.Pop()
		assert(t, item, 0)
		assert(t, err, false)
	})

}
