package day1

import (
	"testing"

	"github.com/adgai19/aoc2022/utils"
)

func TestPart1(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := utils.ReadInputTest()
		want := 24000
		got := Day1a(input)
		utils.Assert(t, want, got)
	})

	t.Run("testing input", func(t *testing.T) {
		input := utils.ReadInput("input.txt")
		want := 75622
		got := Day1a(input)
		utils.Assert(t, want, got)
	})
}

func TestPart2(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := utils.ReadInputTest()
		want := 45000
		got := Day1b(input)
		utils.Assert(t, want, got)
	})

	t.Run("testing input", func(t *testing.T) {
		input := utils.ReadInput("input.txt")
		want := 213159
		got := Day1b(input)
		utils.Assert(t, want, got)
	})
}
