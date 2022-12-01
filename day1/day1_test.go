package day1

import (
	"io/ioutil"
	"testing"

	"github.com/adgai19/utils"
)

func assert(t *testing.T, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("got %v but want %v", got, want)
	}
}
func getDummyInput() string {
	return `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
}
func getInput() string {
	file, _ := ioutil.ReadFile("input.txt")
	input := string(file)
	return input
}

func TestPart1(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := getDummyInput()
		want := 24000
		got := Day1a(input)
		assert(t, want, got)
	})

	t.Run("testing input", func(t *testing.T) {

		input := utils.ReadInput()
		want := 75622
		got := Day1a(input)
		assert(t, want, got)
	})
}
