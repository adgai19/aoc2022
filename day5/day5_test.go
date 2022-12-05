package day4

import (
	"testing"

	"github.com/adgai19/aoc2022/utils"
)

func TestPart1(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := utils.ReadInputTest()
		want := "CMZ"
		got := Day5a(input)
		utils.Assert(t, got, want)
	})
	t.Run("final output", func(t *testing.T) {
		input := utils.ReadInput("input.txt")
		want := "RNZLFZSJH"
		got := Day5a(input)
		utils.Assert(t, want, got)
	})

}

func TestPart2(t *testing.T) {
	t.Run("given test", func(t *testing.T) {
		input := utils.ReadInputTest()
		want := "MCD"
		got := Day5b(input)
		utils.Assert(t, want, got)
	})
	t.Run("final output", func(t *testing.T) {
		input := utils.ReadInput("input.txt")
		want := "CNSFCGJSM"
		got := Day5b(input)
		utils.Assert(t, want, got)
	})
}
