package day1

import (
	"testing"

	"github.com/adgai19/aoc2022/utils"
)

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

func TestPart1(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := getDummyInput()
		want := 24000
		got := Day1a(input)
		utils.Assert(t, want, got)
	})

	t.Run("testing input", func(t *testing.T) {
		input := utils.ReadInput()
		want := 75622
		got := Day1a(input)
		utils.Assert(t, want, got)
	})
}

func TestPart2(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := getDummyInput()
		want := 45000
		got := Day1b(input)
		utils.Assert(t, want, got)
	})

	t.Run("testing input", func(t *testing.T) {
		input := utils.ReadInput()
		want := 213159
		got := Day1b(input)
		utils.Assert(t, want, got)
	})
}
