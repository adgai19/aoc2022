package utils

import (
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	var strs = []string{"peach", "apple", "pear", "plum"}
	want := 2
	got := Index(strs, "pear")
	Assert(t, want, got)

}

func TestIncludes(t *testing.T) {
	var strs = []string{"peach", "apple", "pear", "plum"}
	want := true
	got := Include(strs, "pear")
	Assert(t, want, got)

}
func TestAny(t *testing.T) {
	var strs = []string{"peach", "apple", "pear", "plum"}
	want := true
	got := Any(strs, func(s string) bool { return strings.HasPrefix(s, "p") })
	Assert(t, want, got)

}

func TestAll(t *testing.T) {
	var strs = []string{"peach", "apple", "pear", "plum"}
	want := false
	got := All(strs, func(s string) bool { return strings.HasPrefix(s, "p") })
	Assert(t, want, got)

}

func TestFilter(t *testing.T) {
	var strs = []string{"peach", "apple", "pear", "plum"}
	want := []string{"peach", "pear", "plum"}
	got := Filter(strs, func(s string) bool { return strings.HasPrefix(s, "p") })
	Assert(t, want, got)

}

func TestMap(t *testing.T) {
	var strs = []string{"peach", "apple", "pear", "plum"}
	want := []string{"PEACH", "APPLE", "PEAR", "PLUM"}
	got := Map(strs, strings.ToUpper)
	Assert(t, want, got)

}
