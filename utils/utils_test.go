package utils

import (
	"reflect"
	"testing"
)

func assert[C any](t *testing.T, got, want C) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestReadInput(t *testing.T) {
	want := "foo\n"
	got := ReadInput("input.txt")
	assert(t, want, got)
}

func TestSplitInt(t *testing.T) {
	want := []int{1, 2}
	got := SplitInt("1,2")
	assert(t, want, got)
}

func TestSplitLines(t *testing.T) {
	want := []string{"foo", "bar"}
	got := GetLines(ReadInput("splitLines.txt"))
	assert(t, want, got)
}
func TestSplitComma(t *testing.T) {
	want := []string{"foo", "bar"}
	got := SplitComma("foo,bar")
	assert(t, want, got)
}
