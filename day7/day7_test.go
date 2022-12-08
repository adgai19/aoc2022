package day4

import (
	"testing"

	"github.com/adgai19/aoc2022/utils"
)

func TestPart1(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := utils.ReadInputTest()
		want := 95437
		got := Day7a(input)
		utils.Assert(t, got, want)
	})

	t.Run("input", func(t *testing.T) {
		input := utils.ReadInput("input.txt")
		want := 1449447
		got := Day7a(input)
		utils.Assert(t, got, want)
	})

}

func TestPart2(t *testing.T) {

	t.Run("given test2", func(t *testing.T) {
		input := utils.ReadInputTest()
		want := 24933642
		got := Day7b(input)
		utils.Assert(t, got, want)
	})

	t.Run("given test2", func(t *testing.T) {
		input := utils.ReadInput("input.txt")
		want := 8679207
		got := Day7b(input)
		utils.Assert(t, got, want)
	})
}
