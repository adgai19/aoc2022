package day4

import (
	"testing"

	"github.com/adgai19/aoc2022/utils"
)

func TestPart1(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := utils.ReadInputTest()
		want := 2
		got := Day4a(input)
		utils.Assert(t, want, got)
	})
	t.Run("final output", func(t *testing.T) {
		input := utils.ReadInput("input.txt")
		want := 485
		got := Day4a(input)
		utils.Assert(t, want, got)
	})

}

func TestPart2(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := utils.ReadInputTest()
		want := 4
		got := Day4b(input)
		utils.Assert(t, want, got)
	})
	t.Run("final output", func(t *testing.T) {
		input := utils.ReadInput("input.txt")
		want := 857
		got := Day4b(input)
		utils.Assert(t, want, got)
	})
}
