package utils

import "testing"

func TestQueue(t *testing.T) {
	q := make(Queue[int], 0)
	t.Run("Add items to queue", func(t *testing.T) {
		q.Enque(5)
		want := Queue[int]{5}
		assert(t, q, want)
	})
	t.Run("check if non-empty list is empty", func(t *testing.T) {
		assert(t, q.IsEmpty(), false)
	})
	t.Run("Pop items", func(t *testing.T) {
		item, err := q.Deque()
		assert(t, item, 5)
		assert(t, err, true)
	})
	t.Run("check if list is empty", func(t *testing.T) {
		assert(t, q.IsEmpty(), true)
	})

	t.Run("Pop items from empty list", func(t *testing.T) {
		item, err := q.Deque()
		assert(t, item, 0)
		assert(t, err, false)
	})

}
