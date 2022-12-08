package utils

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func ReadInput(filename string) string {
	file, _ := os.ReadFile(filename)
	input := strings.TrimSuffix(string(file), "\n")

	return input
}

func ReadInputTest() string {
	input := ReadInput("input_test.txt")
	return input
}

func SplitInt(input string) []int {
	split := strings.Split(input, ",")
	out := make([]int, 0)
	for _, val := range split {
		i, _ := strconv.Atoi(val)
		out = append(out, i)
	}
	return out
}

func SplitComma(input string) []string {
	return strings.Split(input, ",")
}

func GetLines(input string) []string {
	out := strings.Split(input, "\n")
	for i, e := range out {
		out[i] = strings.TrimSpace(e)
	}
	return out
}

func Assert[C any](t *testing.T, got, want C) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v but want %v", got, want)
	}
}
