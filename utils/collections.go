package utils

func Index[T comparable](vs []T, t T) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include returns true if the target string t is in the slice.

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any returns true if one of the strings in the slice satisfies the predicate f.

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all of the strings in the slice satisfy the predicate f.

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all strings in the slice that satisfy the predicate f.

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying the function f to each string in the original slice.

func Map[T any, V any](vs []T, f func(T) T) []T {
	vsm := make([]T, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
