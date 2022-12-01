package utils

import (
	"reflect"
	"testing"
)

func assertString(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("got %v but want %v", got, want)
	}
}

func assertArrays(t *testing.T, want, got []int) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v but want %v", got, want)
	}
}
func TestReadInput(t *testing.T) {
	want := "foo\n"
	got := ReadInput()
	assertString(t, want, got)
}
func TestSplitInt(t *testing.T) {
	want := []int{1, 2}
	got := SplitInt("1,2")
	assertArrays(t, want, got)

}
